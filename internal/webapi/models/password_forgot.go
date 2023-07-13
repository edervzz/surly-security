package models

// @Description New Password forgot flow
type NewPasswordForgot struct {
	Email string `json:"email" mapper:"email"`
} //@name NewPasswordForgot

// @Description Password forgot flow
type PasswordForgot struct {
	Token string `json:"token" example:"2eba30ff-adb8-478b-913f-ace363acbd34" mapper:"token"`
} //@name PasswordForgot
