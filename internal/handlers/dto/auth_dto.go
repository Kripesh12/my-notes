package dto

import "github.com/kripesh12/my-notes/internal/validator"

type RegisterResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *RegisterRequest) Validate() error {
	if err := validator.ValidateEmail(r.Email); err != nil {
		return err
	}

	if err := validator.ValidatePassword(r.Password); err != nil {
		return err
	}

	return nil
}
