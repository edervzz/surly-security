package messages

type ForgotPasswordResponse struct {
	Token string `mapper:"token"`
}
