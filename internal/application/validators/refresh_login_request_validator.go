package validators

import (
	"surly-security/internal/application/messages"
	"surly-security/internal/application/resources"
	"surly-security/toolkit/localizer"

	"github.com/edervzz/mistake"
)

func RefreshLoginRequestValidator(request messages.RefreshLoginRequest, localizer localizer.ILocalizer) *mistake.M {
	if request.RefreshToken == "" {
		return mistake.New(
			resources.VALIDATION,
			resources.REFRESH_LOGIN_TOKEN,
			localizer.Localize(resources.REFRESH_LOGIN_TOKEN),
		)
	}
	return nil
}
