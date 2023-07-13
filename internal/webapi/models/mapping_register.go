package models

import (
	"fmt"
	"surly-security/internal/application/messages"

	"github.com/devfeel/mapper"
)

func init() {
	mapper.Register(&NewSignUp{})
	mapper.Register(&messages.SignupRequest{})

	mapper.Register(&SignUp{})
	mapper.Register(&messages.SignupResponse{})

	mapper.Register(&NewLogin{})
	mapper.Register(&messages.LoginRequest{})

	mapper.Register(&Login{})
	mapper.Register(&messages.LoginResponse{})

	mapper.Register(&NewRefreshLogin{})
	mapper.Register(&messages.RefreshLoginRequest{})

	mapper.Register(&RefreshLogin{})
	mapper.Register(&messages.RefreshLoginResponse{})

	mapper.Register(&NewChangePassword{})
	mapper.Register(&messages.ChangePasswordRequest{})

	mapper.Register(&ChangePassword{})
	mapper.Register(&messages.ChangePasswordResponse{})

	mapper.Register(&NewPasswordForgot{})
	mapper.Register(&messages.ForgotPasswordRequest{})

	mapper.Register(&PasswordForgot{})
	mapper.Register(&messages.ForgotPasswordResponse{})

}

func Map[T interface{}](source interface{}) T {
	destination := new(T)
	err := mapper.Mapper(source, destination)
	if err != nil {
		panic(fmt.Sprintf("mapping.Map: %s", err.Error()))
	}
	return *destination
}

func Adapt(source interface{}, destination interface{}) {
	err := mapper.Mapper(source, destination)
	if err != nil {
		panic(fmt.Sprintf("mapping.Map: %s", err.Error()))
	}
}
