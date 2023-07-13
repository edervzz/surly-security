package endpoints

import (
	"encoding/json"
	"net/http"
	"surly-security/internal/application/commands"
	"surly-security/internal/application/messages"
	"surly-security/internal/domain/ports"
	"surly-security/internal/webapi/models"
	"surly-security/toolkit/localizer"
	"surly-security/toolkit/services"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

// Confirm sign up godoc
//
//	@Summary		Confirm sign up
//	@Description	Confirm user created
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			token  path    string  true    "Confirmation token"
//	@Success		200 {object} models.EmailConfirmToken
//	@Failure		400 "Invalid token"
//	@Failure		500
//	@Router			/users/confirm/email/{token} [post]
func ConfirmUserByEmail(w http.ResponseWriter, r *http.Request) {
	request := messages.UserConfirmEmailRequest{Token: chi.URLParam(r, "token")}
	_, mt := commands.UserConfirmEmailHandler(
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

	result := models.EmailConfirmToken{
		IsSuccess: true,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
