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

func RefreshLoginHandler(
	ctx context.Context,
	uow ports.IUnitOfWork,
	localizer localizer.ILocalizer,
	logger *zap.Logger,
	request messages.RefreshLoginRequest,
) (messages.RefreshLoginResponse, *mistake.M) {
	// 1. request validation
	if mt := validators.RefreshLoginRequestValidator(request, localizer); mt != nil {
		return messages.RefreshLoginResponse{}, mt
	}
	// 2. business rules validation
	if mt := validators.RefreshLoginRequestBizValidator(ctx, &request, localizer, uow); mt != nil {
		return messages.RefreshLoginResponse{}, mt
	}
	// 3. create tokens
	accessToken, refreshToken, mt := extensions.GenerateTokens(ctx,
		request.User, request.User.Email, request.User.ID, entities.ExpiresAt_20min, entities.ExpiresAt_20days, logger, localizer, uow)

	if mt != nil {
		return messages.RefreshLoginResponse{}, mt
	}

	// 4. deactivate token used
	request.TokenUser.IsActive = false
	if err := uow.BeginTransaction(ctx); err == nil {
		if err = uow.TokenUser().Update(ctx, &request.TokenUser); err == nil {
			err = uow.CommitWork()
		}
	}

	return messages.RefreshLoginResponse{
		AccessToken:  accessToken,
		TokenType:    entities.Bearer,
		ExpiresIn:    entities.ExpiresAt_20min,
		RefreshToken: refreshToken,
	}, nil

}
