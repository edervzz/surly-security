package models

// @Description Change current password
type NewChangePassword struct {
	CurrentPassword string `json:"current_password" mapper:"current_password"`
	// "One upper, one lower, one number, one special @#$%&, len: 8-16"
	NewPassword string `json:"new_password" validate:"required" minLength:"8" maxLength:"16" format:"password" mapper:"new_password"`
} //@name NewChangePassword

// @Description Result of change current password
type ChangePassword struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." mapper:"access_token"`
	TokenType    string `json:"token_type" example:"Bearer" mapper:"token_type"`
	ExpiresIn    int    `json:"expires_in" example:"1800" mapper:"expires_in"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." mapper:"refresh_token"`
} //@name ChangePassword
