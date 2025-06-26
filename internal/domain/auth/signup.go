package auth

import (
	"goapp/pkg/validation"
)

type SignupDTO struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (c *SignupDTO) Validate() error {
	return validation.ValidateStruct(c)
}

func Signup(dto *SignupDTO) (profileID int, err error) {
	return
}
