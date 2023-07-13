package entities

import "github.com/edervzz/maya"

type User struct {
	maya.IEntity `tname:"users"`
	maya.IAuditable
	ID             int64  `tcol:"id" pk:"true" auto:"true"`
	Email          string `tcol:"email"`
	Fullname       string `tcol:"fullname"`
	PasswordHash   string `tcol:"password_hash"`
	EmailConfirmed bool   `tcol:"email_confirmed"`
	PhoneConfirmed bool   `tcol:"phone_confirmed"`
	IsLocked       bool   `tcol:"is_locked"`
	IsActive       bool   `tcol:"is_active"`
	Intents        int    `tcol:"intents"`
}
