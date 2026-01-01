package validator

import (
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
var (
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex = regexp.MustCompile(`[a-z]`)
	digitRegex     = regexp.MustCompile(`\d`)
	specialRegex   = regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{}|;:'",.<>/?]`)
)

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}

	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if !uppercaseRegex.MatchString(password) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !lowercaseRegex.MatchString(password) {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !digitRegex.MatchString(password) {
		return errors.New("password must contain at least one number")
	}
	if !specialRegex.MatchString(password) {
		return errors.New("password must contain at least one special character")
	}
	return nil
}
