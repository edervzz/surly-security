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

// Login godoc
//
//	@Summary		Login
//	@Description	Login session
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			RequestBody     body    models.NewLogin    true    "Login User Information"
//	@Success		200 {object} models.Login
//	@Failure		400 "Login failed"
//	@Failure		500
//	@Router			/users/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	newLogin := models.NewLogin{}
	if err := json.NewDecoder(r.Body).Decode(&newLogin); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	request := models.Map[messages.LoginRequest](&newLogin)

	response, mt := commands.LoginHandler(
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

	result := models.Map[models.Login](&response)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
