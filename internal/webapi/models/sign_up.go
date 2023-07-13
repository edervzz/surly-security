package models

// @Description User account information
type NewSignUp struct {
	Email    string `json:"email" validate:"required" example:"johndoe@mail.com" mapper:"email"`
	Fullname string `json:"fullname" validate:"required" minLength:"5" maxLength:"30" example:"John Doe" mapper:"fullname"`
	// "One upper, one lower, one number, one special @#$%&, len: 8-16 "
	Password string `json:"password" validate:"required" minLength:"8" maxLength:"16" format:"password" mapper:"password"`
} //@name NewSignUp

// @Description Confirmation token
type SignUp struct {
	Email        string `json:"email" example:"johndoe@mail.com" mapper:"email"`
	ConfirmToken string `json:"confirm_token" example:"2eba30ff-adb8-478b-913f-ace363acbd34" mapper:"confirm_token"`
} //@name SignUp
