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

	"go.uber.org/zap"
)

// ForgotPassword godoc
//
//	@Summary		Forgot password flow
//	@Description	Forgot password flow
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			RequestBody     body    models.NewPasswordForgot    true    "New Forgot Password info"
//	@Success		200 {object} models.PasswordForgot
//	@Failure		400 "Something is wrong"
//	@Failure		500
//	@Router			/users/password/forgot [post]
func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	newForgotPassword := models.NewPasswordForgot{}
	err := json.NewDecoder(r.Body).Decode(&newForgotPassword)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	request := models.Map[messages.ForgotPasswordRequest](&newForgotPassword)

	response, mt := commands.ForgotPasswordHandler(
		r.Context(),
		services.Get[ports.IUnitOfWork](),
		services.Get[localizer.ILocalizer](),
		services.Get[*zap.Logger](),
		request,
	)

	// 4. on error, return
	if mt != nil {
		w.WriteHeader(mt.ReturnCode)
		json.NewEncoder(w).Encode(mt)
		return
	}

	// 5. message map to model
	result := models.Map[models.PasswordForgot](&response)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
