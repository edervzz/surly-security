package models

// @Description Reset password forgot
type NewResetPasswordForgot struct {
	ResetPasswordToken string `json:"reset_password_token" example:"2eba30ff-adb8-478b-913f-ace363acbd34" mapper:"reset_password_token"`
	// "One upper, one lower, one number, one special @#$%&, len: 8-16"
	NewPassword string `json:"new_password" validate:"required" minLength:"8" maxLength:"16" format:"password" mapper:"new_password"`
} //@name NewResetPasswordForgot

// @Description Result Reset password forgot
type ResetPasswordForgot struct {
	IsSuccess bool
} //@name ResetPasswordForgot
