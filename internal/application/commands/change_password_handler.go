package commands

import (
	"context"
	"surly-security/internal/application/extensions"
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/internal/application/validators"
	"surly-security/internal/domain/entities"
	"surly-security/internal/domain/ports"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"
	"github.com/go-chi/jwtauth"
	"golang.org/x/crypto/bcrypt"

	"go.uber.org/zap"
)

func ChangePasswordHandler(ctx context.Context,
	uow ports.IUnitOfWork,
	localizer localizer.ILocalizer,
	logger *zap.Logger,
	request messages.ChangePasswordRequest,
) (messages.ChangePasswordResponse, *mistake.M) {
	// 1. request validated
	_, claims, _ := jwtauth.FromContext(ctx)

	if user, ok := claims["sub"].(string); ok {
		request.UserEmail = user
	}

	if mt := validators.ChangePasswordRequestValidator(request, localizer); mt != nil {
		return messages.ChangePasswordResponse{}, mt
	}
	logger.Debug("request validate")
	// 2. business rules validated
	if mt := validators.ChangePasswordRequestBizValidator(ctx, &request, uow, localizer); mt != nil {
		return messages.ChangePasswordResponse{}, mt
	}
	logger.Debug("business rules validated")
	// 3. prepare user entity
	passwordHashed, _ := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	request.User.PasswordHash = string(passwordHashed)
	logger.Debug("entity prepared")
	// 4. save entity
	defer uow.RollbackWork()
	err := *new(error)
	if err = uow.BeginTransaction(ctx); err == nil {
		if err = uow.User().Update(ctx, &request.User); err == nil {
			err = uow.CommitWork()
		}
	}
	// 4.1 handle error
	if err != nil {
		return messages.ChangePasswordResponse{}, mistake.New(
			resources.BAD_REQUEST,
			resources.DB,
			err.Error(),
		)
	}

	accessToken, refreshToken, mt := extensions.GenerateTokens(ctx,
		request.User,
		request.User.Email,
		request.User.ID,
		entities.ExpiresAt_20min,
		entities.ExpiresAt_20days,
		logger,
		localizer,
		uow)

	if mt != nil {
		return messages.ChangePasswordResponse{}, mt
	}

	return messages.ChangePasswordResponse{
		AccessToken:  accessToken,
		TokenType:    entities.Bearer,
		ExpiresIn:    entities.ExpiresAt_20min,
		RefreshToken: refreshToken,
	}, nil

}
