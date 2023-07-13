package commands

import (
	"context"
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/internal/application/validators"
	"surly-security/internal/domain/entities"
	"surly-security/internal/domain/ports"
	"surly-security/toolkit/localizer"
	"time"

	"github.com/edervzz/mistake"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(
	ctx context.Context,
	uow ports.IUnitOfWork,
	logger *zap.Logger,
	localizer localizer.ILocalizer,
	request messages.SignupRequest,
) (messages.SignupResponse, *mistake.M) {
	// 1. request validations
	mt := validators.SignUpRequestValidator(request, localizer)
	if mt != nil {
		return messages.SignupResponse{}, mt
	}
	logger.Debug("request validated")
	// 2. business validations
	mt = validators.SignUpRequestBizValidator(request, uow, localizer, ctx)
	if mt != nil {
		return messages.SignupResponse{}, mt
	}
	logger.Debug("business rules validated")
	// 3. prepare user entity
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := entities.User{
		Email:          request.Email,
		Fullname:       request.Fullname,
		PasswordHash:   string(hashedPassword),
		EmailConfirmed: false,
		PhoneConfirmed: false,
		IsLocked:       false,
		IsActive:       false,
		Intents:        0,
	}
	// 3.1 prepare user tokens entity
	userToken := entities.TokenUser{
		Token:     uuid.NewString(),
		Type:      entities.TokenTypeConfirmEmail,
		UserID:    0,
		ExpiresAt: time.Now().UTC().Add(time.Second * entities.ExpiresAt_20min).Format(time.DateTime),
		IsActive:  true,
	}
	logger.Debug("entities prepared")
	// 4. save entity
	defer uow.RollbackWork()
	err := *new(error)
	if err = uow.BeginTransaction(ctx); err == nil {
		if err = uow.User().Create(ctx, &user); err == nil {
			userToken.UserID = user.ID
			if err = uow.TokenUser().Create(ctx, &userToken); err == nil {
				err = uow.CommitWork()
			}
		}
	}
	// 4.1 handle error
	if err != nil {
		return messages.SignupResponse{}, mistake.New(
			resources.BAD_REQUEST,
			resources.DB,
			err.Error(),
		)
	}

	return messages.SignupResponse{
		Email:        user.Email,
		ConfirmToken: userToken.Token,
	}, nil
}
