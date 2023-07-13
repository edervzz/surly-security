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

func UserConfirmEmailRequestBizValidator(ctx context.Context, uow ports.IUnitOfWork, request *messages.UserConfirmEmailRequest, localizer localizer.ILocalizer) *mistake.M {
	// get token
	tokenUser, err := uow.TokenUser().Read(request.Token, ctx)
	if err != nil {
		return mistake.New(
			resources.SERVER_ERROR,
			resources.DB,
			err.Error())
	}
	// token is required
	if tokenUser == (entities.TokenUser{}) {
		return mistake.New(
			resources.VALIDATION,
			resources.USER_CONFIRM_TOKEN,
			localizer.Localize(resources.USER_CONFIRM_TOKEN))
	}
	// check expiration
	expiresAt, err := time.Parse(time.DateTime, tokenUser.ExpiresAt)
	if err != nil {
		return mistake.New(
			resources.SERVER_ERROR,
			resources.INTERNAL,
			err.Error())
	}
	currentTime := time.Now().UTC()
	if expiresAt.Before(currentTime) {
		return mistake.New(
			resources.VALIDATION,
			resources.USER_CONFIRM_TOKEN_VALID,
			localizer.Localize(resources.USER_CONFIRM_TOKEN_VALID))
	}
	// check is an active token
	if !tokenUser.IsActive {
		return mistake.New(
			resources.VALIDATION,
			resources.USER_CONFIRM_TOKEN_ACTIVE,
			localizer.Localize(resources.USER_CONFIRM_TOKEN_ACTIVE))
	}
	// get user
	user, err := uow.User().Read(tokenUser.UserID, ctx)
	if err != nil {
		return mistake.New(
			resources.SERVER_ERROR,
			resources.DB,
			err.Error())
	}
	// user exists
	if user == (entities.User{}) {
		return mistake.New(
			resources.NOT_FOUND,
			resources.USER_CONFIRM_EXIST,
			localizer.Localize(resources.USER_CONFIRM_EXIST))
	}
	// user inactive
	if user.IsActive {
		return mistake.New(
			resources.VALIDATION,
			resources.USER_INACTIVE,
			localizer.Localize(resources.USER_INACTIVE))
	}

	request.User = user
	request.TokenUser = tokenUser
	return nil
}
