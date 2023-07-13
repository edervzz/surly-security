package validators

import (
	"context"
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/internal/domain/entities"
	"surly-security/internal/domain/ports"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"
	"golang.org/x/crypto/bcrypt"
)

func ChangePasswordRequestBizValidator(ctx context.Context, request *messages.ChangePasswordRequest, uow ports.IUnitOfWork, localizer localizer.ILocalizer) *mistake.M {
	user, err := uow.User().ReadByExternalID(ctx, request.UserEmail)
	if err != nil {
		return mistake.New(
			resources.BAD_REQUEST,
			resources.DB,
			err.Error(),
		)
	}
	// user must exist
	if user == (entities.User{}) {
		return mistake.New(
			resources.VALIDATION,
			resources.LOGIN_USER_PASSWORD,
			localizer.Localize(resources.LOGIN_USER_PASSWORD),
		)
	}
	// password must match
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.CurrentPassword)); err != nil {
		return mistake.New(
			resources.VALIDATION,
			resources.LOGIN_USER_PASSWORD,
			localizer.Localize(resources.LOGIN_USER_PASSWORD),
		)
	}
	// user must be active
	if !user.IsActive {
		return mistake.New(
			resources.VALIDATION,
			resources.LOGIN_USER_ACTIVE,
			localizer.Localize(resources.LOGIN_USER_ACTIVE),
		)
	}
	// user must be without locks
	if user.IsLocked {
		return mistake.New(
			resources.VALIDATION,
			resources.LOGIN_USER_LOCK,
			localizer.Localize(resources.LOGIN_USER_LOCK),
		)
	}

	request.User = user

	return nil
}
