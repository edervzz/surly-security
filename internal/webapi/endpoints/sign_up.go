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

// SignUp godoc
//
//	@Summary		Sign Up
//	@Description	Register a new user with password
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			RequestBody     body    models.NewSignUp    true    "User Information"
//	@Success		200 {object} models.SignUp
//	@Failure		400
//	@Failure		409  "User already exists."
//	@Failure		500
//	@Router			/users/sign-up [post]
func SignUp(w http.ResponseWriter, r *http.Request) {
	newSignup := models.NewSignUp{}
	err := json.NewDecoder(r.Body).Decode(&newSignup)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	request := models.Map[messages.SignupRequest](&newSignup)

	response, mt := commands.SignUpHandler(
		r.Context(),
		services.Get[ports.IUnitOfWork](),
		services.Get[*zap.Logger](),
		services.Get[localizer.ILocalizer](),
		request,
	)

	if mt != nil {
		w.WriteHeader(mt.ReturnCode)
		json.NewEncoder(w).Encode(mt)
		return
	}

	result := models.Map[models.SignUp](&response)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
