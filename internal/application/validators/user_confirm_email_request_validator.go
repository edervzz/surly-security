package validators

import (
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"

	"github.com/go-playground/validator"
)

func UserConfirmEmailRequestValidator(request messages.UserConfirmEmailRequest, localizer localizer.ILocalizer) *mistake.M {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return mistake.New(
			resources.VALIDATION,
			resources.USER_CONFIRM_TOKEN_EMPTY,
			localizer.Localize(resources.USER_CONFIRM_TOKEN_EMPTY),
		)
	}
	return nil
}
