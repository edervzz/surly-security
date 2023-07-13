package models

// @Description New Refresh Login
type NewLoginRefresh struct {
	RefreshToken string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." mapper:"refresh_token"`
} //@name NewLoginRefresh

// @Description Refresh Login
type LoginRefresh struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." mapper:"access_token"`
	TokenType    string `json:"token_type" example:"Bearer" mapper:"token_type"`
	ExpiresIn    int    `json:"expires_in" example:"1800" mapper:"expires_in"`
	RefreshToken string `json:"refresh_token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." mapper:"refresh_token"`
} //@name LoginRefresh
