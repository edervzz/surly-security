package validators

import (
	"context"
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/internal/domain/entities"
	"surly-security/internal/domain/ports"
	"surly-security/toolkit/localizer"
	"time"

	"github.com/edervzz/mistake"
)

func RefreshLoginRequestBizValidator(ctx context.Context, request *messages.RefreshLoginRequest, localizer localizer.ILocalizer, uow ports.IUnitOfWork) *mistake.M {
	tokenUser, err := uow.TokenUser().Read(request.RefreshToken, ctx)
	if err != nil {
		return mistake.New(
			resources.BAD_REQUEST,
			resources.DB,
			err.Error(),
		)

	}
	// token user must exists
	if tokenUser == (entities.TokenUser{}) {
		return mistake.New(
			resources.VALIDATION,
			resources.REFRESH_LOGIN_TOKEN_VALID,
			localizer.Localize(resources.REFRESH_LOGIN_TOKEN_VALID),
		)
	}
	// token must be active
	if !tokenUser.IsActive {
		return mistake.New(
			resources.VALIDATION,
			resources.REFRESH_LOGIN_TOKEN_VALID,
			localizer.Localize(resources.REFRESH_LOGIN_TOKEN_VALID),
		)
	}

	expiresAt, err := time.Parse(time.DateTime, tokenUser.ExpiresAt)
	if err != nil {
		return mistake.New(
			resources.SERVER_ERROR,
			resources.INTERNAL,
			err.Error())
	}
	// token must be valid
	currentTime := time.Now().UTC()
	if expiresAt.Before(currentTime) {
		return mistake.New(
			resources.VALIDATION,
			resources.REFRESH_LOGIN_TOKEN_VALID,
			localizer.Localize(resources.REFRESH_LOGIN_TOKEN_VALID),
		)
	}

	user, err := uow.User().Read(tokenUser.UserID, ctx)
	if err != nil {
		return mistake.New(
			resources.SERVER_ERROR,
			resources.INTERNAL,
			err.Error())
	}

	request.User = user
	request.TokenUser = tokenUser
	return nil
}
