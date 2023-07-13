package entities

import "github.com/edervzz/maya"

type TokenUser struct {
	maya.IEntity `tname:"token_users"`
	maya.IAuditable
	Token     string `tcol:"token" pk:"true"`
	Type      string `tcol:"type"`
	UserID    int64  `tcol:"user_id"`
	ExpiresAt string `tcol:"expires_at"`
	IsActive  bool   `tcol:"is_active"`
}

const (
	Bearer                = "Bearer"
	TokenTypeAccess       = "access"
	TokenTypeRefresh      = "refresh"
	TokenTypeResetPwd     = "resetPassword"
	TokenTypeConfirmEmail = "confirmEmail"
)

const (
	ExpiresAt_20min        = 1200
	ExpiresAt_90min        = 5400
	ExpiresAt_20days int64 = 1728000
)
