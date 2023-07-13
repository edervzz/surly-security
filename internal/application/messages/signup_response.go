package messages

type SignupResponse struct {
	Email        string `mapper:"email"`
	ConfirmToken string `mapper:"confirm_token"`
}
