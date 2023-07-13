package messages

import "surly-security/internal/domain/entities"

type LoginRequest struct {
	Username string `mapper:"username" validator:"required"`
	Password string `mapper:"password" validator:"required"`
	User     entities.User
}
