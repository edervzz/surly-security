package messages

import "surly-security/internal/domain/entities"

type ChangePasswordRequest struct {
	CurrentPassword string `mapper:"current_password"`
	NewPassword     string `mapper:"new_password" validate:"password"`
	UserEmail       string

	User entities.User
}
