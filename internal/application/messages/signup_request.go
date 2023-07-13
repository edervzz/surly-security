package messages

import "surly-security/internal/domain/entities"

type SignupRequest struct {
	Email    string `validate:"required,email" mapper:"email"`
	Fullname string `validate:"required" mapper:"fullname"`
	Password string `validate:"required" mapper:"password"`

	User entities.User
}
