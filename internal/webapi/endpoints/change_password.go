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

// ChangePassword godoc
//
//	@Summary		Change password
//	@Description	Change current password including a valid access token.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			RequestBody     body    models.NewChangePassword    true    "Change password Information"
//	@Success		200 {object} models.ChangePassword
//	@Failure		400 "Login failed"
//	@Failure		500
//	@Router			/users/me/password [put]
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	// 1. body decode to model
	newChangePassword := models.NewChangePassword{}
	if err := json.NewDecoder(r.Body).Decode(&newChangePassword); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	// 2. model map to message
	request := models.Map[messages.ChangePasswordRequest](&newChangePassword)
	// 3. call command
	response, mt := commands.ChangePasswordHandler(
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
	result := models.Map[models.ChangePassword](&response)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
