package validators

import (
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"
	"github.com/go-playground/validator"
)

func ChangePasswordRequestValidator(request messages.ChangePasswordRequest, localizer localizer.ILocalizer) *mistake.M {
	v := validator.New()
	v.RegisterValidation("password", PasswordValidator)
	err := v.Struct(request)
	if err != nil {
		return mistake.NewStructValidation(
			err,
			resources.VALIDATION,
			resources.PASSWORD_CHECK,
			localizer.Localize(resources.PASSWORD_CHECK),
		)
	}

	if request.CurrentPassword == "" {
		return mistake.New(
			resources.VALIDATION,
			resources.CHANGE_PASS_CURRENT,
			localizer.Localize(resources.CHANGE_PASS_CURRENT),
		)
	}

	return nil
}
