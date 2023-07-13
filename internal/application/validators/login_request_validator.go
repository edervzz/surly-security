package validators

import (
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"

	"github.com/go-playground/validator"
)

func LoginRequestValidator(request messages.LoginRequest, localizer localizer.ILocalizer) *mistake.M {
	v := validator.New()
	err := v.Struct(request)
	if err != nil {
		return mistake.NewStructValidation(
			err,
			resources.VALIDATION,
			resources.LOGIN_USER_REQUEST,
			localizer.Localize(resources.LOGIN_USER_REQUEST),
		)
	}
	return nil
}
