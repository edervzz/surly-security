package endpoints

import (
	"encoding/json"
	"net/http"
	"net/url"
	"surly-security/internal/application/commands"
	"surly-security/internal/application/messages"
	"surly-security/internal/domain/ports"
	"surly-security/internal/webapi/models"
	"surly-security/toolkit/localizer"
	"surly-security/toolkit/services"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// Login godoc
//
//	@Summary		Refresh Login
//	@Description	Refresh login session
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			token  path    string  true    "Refresh token"
//	@Success		200 {object} models.RefreshLogin
//	@Failure		400 "Login failed"
//	@Failure		500
//	@Router			/users/login/refresh/{token} [post]
func RefreshLogin(w http.ResponseWriter, r *http.Request) {
	token, _ := url.PathUnescape(chi.URLParam(r, "token"))
	newLoginRefresh := models.NewLoginRefresh{
		RefreshToken: token,
	}

	request := models.Map[messages.RefreshLoginRequest](&newLoginRefresh)

	response, mt := commands.RefreshLoginHandler(
		r.Context(),
		services.Get[ports.IUnitOfWork](),
		services.Get[localizer.ILocalizer](),
		services.Get[*zap.Logger](),
		request,
	)

	if mt != nil {
		w.WriteHeader(mt.ReturnCode)
		json.NewEncoder(w).Encode(mt)
		return
	}

	result := models.Map[models.RefreshLogin](&response)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
