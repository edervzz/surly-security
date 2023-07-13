package commands

import (
	"context"
	"surly-security/internal/application/extensions"
	"surly-security/internal/application/messages"
	"surly-security/internal/application/validators"
	"surly-security/internal/domain/entities"
	"surly-security/internal/domain/ports"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"

	"go.uber.org/zap"
)

func LoginHandler(
	ctx context.Context,
	uow ports.IUnitOfWork,
	localizer localizer.ILocalizer,
	logger *zap.Logger,
	request messages.LoginRequest,
) (messages.LoginResponse, *mistake.M) {
	// 1. request validation
	if mt := validators.LoginRequestValidator(request, localizer); mt != nil {
		return messages.LoginResponse{}, mt
	}
	logger.Debug("request validated")
	// 2. request business validation
	if mt := validators.LoginRequestBizValidator(ctx, &request, uow, localizer); mt != nil {
		return messages.LoginResponse{}, mt
	}
	logger.Debug("business rules validated")
	// 3. create tokens
	accessToken, refreshToken, mt := extensions.GenerateTokens(ctx,
		request.User, request.User.Email, request.User.ID, entities.ExpiresAt_90min, entities.ExpiresAt_20days, logger, localizer, uow)

	if mt != nil {
		return messages.LoginResponse{}, mt
	}

	return messages.LoginResponse{
		AccessToken:  accessToken,
		TokenType:    entities.Bearer,
		ExpiresIn:    entities.ExpiresAt_20min,
		RefreshToken: refreshToken,
	}, nil

}
