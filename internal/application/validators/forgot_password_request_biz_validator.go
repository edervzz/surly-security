package validators

import (
	"context"
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/internal/domain/entities"
	"surly-security/internal/domain/ports"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"
)

func ForgotPasswordRequestBizValidator(ctx context.Context, request *messages.ForgotPasswordRequest, uow ports.IUnitOfWork, localizer localizer.ILocalizer) *mistake.M {
	user, err := uow.User().ReadByExternalID(ctx, request.Email)
	if err != nil {
		return mistake.New(
			resources.SERVER_ERROR,
			resources.DB,
			err.Error())
	}

	if user == (entities.User{}) {
		return mistake.New(
			resources.VALIDATION,
			resources.FORGOT_PASSWORD_USER,
			localizer.Localize(resources.FORGOT_PASSWORD_USER))
	}

	if !user.IsActive {
		return mistake.New(
			resources.VALIDATION,
			resources.FORGOT_PASSWORD_ACTIVE,
			localizer.Localize(resources.FORGOT_PASSWORD_ACTIVE))
	}

	request.User = user

	return nil
}
