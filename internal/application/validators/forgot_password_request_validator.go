package validators

import (
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"
	"github.com/go-playground/validator"
)

func ForgotPasswordRequestValidator(request messages.ForgotPasswordRequest, localizer localizer.ILocalizer) *mistake.M {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return mistake.NewStructValidation(
			err,
			resources.VALIDATION,
			resources.FORGOT_PASSWORD_STRUCT,
			localizer.Localize(resources.FORGOT_PASSWORD_STRUCT),
		)
	}
	return nil
}
