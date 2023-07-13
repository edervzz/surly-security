package validators

import (
	"regexp"

	"github.com/go-playground/validator"
)

func PasswordValidator(fl validator.FieldLevel) bool {
	// Minimum eight characters, at least one letter, one number and one special character
	tests := []string{".{7,}", "[a-z]", "[A-Z]", "[0-9]", "[!@#$&*]"}
	for _, test := range tests {
		t, _ := regexp.MatchString(test, fl.Field().String())
		if !t {
			return false
		}
	}
	return true
}
