package validate

import (
	"github.com/golang-test-task/pkg/common/errors"
	"strings"
)

var invalidEmail = errors.New("invalid email")

func ValidateEmail(email string) (string, error) {
	if !IsEmail(email) {
		return "", invalidEmail
	}
	email = strings.ToLower(email)
	return email, nil
}
