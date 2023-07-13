package messages

import "surly-security/internal/domain/entities"

type UserConfirmEmailRequest struct {
	Token     string
	User      entities.User
	TokenUser entities.TokenUser
}
