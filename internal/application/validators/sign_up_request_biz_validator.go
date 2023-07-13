package validators

import (
	"context"
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/internal/domain/entities"
	"surly-security/internal/domain/ports"

	"github.com/edervzz/mistake"

	"surly-security/toolkit/localizer"
)

func SignUpRequestBizValidator(request messages.SignupRequest, uow ports.IUnitOfWork, localizer localizer.ILocalizer, ctx context.Context) *mistake.M {
	user, err := uow.User().ReadByExternalID(ctx, request.Email)
	if err != nil {
		return mistake.New(
			resources.BAD_REQUEST,
			resources.DB,
			err.Error(),
		)

	}
	if user != (entities.User{}) {
		mistake.AppendValue(request.Email)
		return mistake.New(
			resources.DUPLICATED,
			resources.CREATE_USER_EXIST,
			mistake.Formatter(localizer.Localize(resources.CREATE_USER_EXIST)),
		)
	}

	return nil
}
