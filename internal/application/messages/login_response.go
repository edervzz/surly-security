package messages

type LoginResponse struct {
	AccessToken  string `mapper:"access_token"`
	TokenType    string `mapper:"token_type"`
	ExpiresIn    int    `mapper:"expires_in"`
	RefreshToken string `mapper:"refresh_token"`
}
