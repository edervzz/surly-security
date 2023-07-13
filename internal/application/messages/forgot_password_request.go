package messages

import "surly-security/internal/domain/entities"

type ForgotPasswordRequest struct {
	Email string `mapper:"email" validate:"required"`
	User  entities.User
}
