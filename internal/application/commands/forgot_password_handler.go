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
)

func ForgotPasswordHandler(
	ctx context.Context,
	uow ports.IUnitOfWork,
	localizer localizer.ILocalizer,
	logger *zap.Logger,
	request messages.ForgotPasswordRequest,
) (messages.ForgotPasswordResponse, *mistake.M) {
	// 1. request validation
	if mt := validators.ForgotPasswordRequestValidator(request, localizer); mt != nil {
		return *new(messages.ForgotPasswordResponse), mt
	}
	logger.Debug("request validated")
	// 2. business validation
	if mt := validators.ForgotPasswordRequestBizValidator(ctx, &request, uow, localizer); mt != nil {
		return *new(messages.ForgotPasswordResponse), mt
	}
	logger.Debug("business rules valudated")
	// 3. prepare entity
	userToken := entities.TokenUser{
		Token:     uuid.NewString(),
		Type:      entities.TokenTypeResetPwd,
		UserID:    request.User.ID,
		ExpiresAt: time.Now().UTC().Add(time.Second * entities.ExpiresAt_20min).Format(time.DateTime),
		IsActive:  true,
	}
	// 4. save entity
	defer uow.RollbackWork()
	err := *new(error)
	if err = uow.BeginTransaction(ctx); err == nil {
		if err = uow.TokenUser().Create(ctx, &userToken); err == nil {
			err = uow.CommitWork()
		}
	}
	// 4.1 handle error
	if err != nil {
		return messages.ForgotPasswordResponse{}, mistake.New(
			resources.BAD_REQUEST,
			resources.DB,
			err.Error(),
		)
	}

	return messages.ForgotPasswordResponse{
		Token: userToken.Token,
	}, nil
}
