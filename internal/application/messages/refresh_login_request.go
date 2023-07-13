package messages

import "surly-security/internal/domain/entities"

type RefreshLoginRequest struct {
	RefreshToken string `mapper:"refresh_token"`
	User         entities.User
	TokenUser    entities.TokenUser
}
