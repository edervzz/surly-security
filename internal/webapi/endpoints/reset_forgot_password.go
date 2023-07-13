package endpoints

import (
	"encoding/json"
	"net/http"
	"surly-security/internal/webapi/models"
)

// ResetForgotPassword godoc
//
//	@Summary		Forgot password flow
//	@Description	Forgot password flow
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			RequestBody  body    models.NewResetPasswordForgot  true    "Reset password information"
//	@Success		200 {object} models.ResetPasswordForgot
//	@Failure		400 "Something is wrong"
//	@Failure		500
//	@Router			/users/password/reset [post]
func ResetForgotPassword(w http.ResponseWriter, r *http.Request) {
	newResetForgPassword := models.NewResetPasswordForgot{}
	err := json.NewDecoder(r.Body).Decode(&newResetForgPassword)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	result := models.ResetPasswordForgot{
		IsSuccess: true,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
