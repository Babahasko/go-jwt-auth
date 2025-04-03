package req

import (
	"github.com/Babahasko/go-jwt-auth/internal/validation"
	"github.com/go-playground/validator/v10"
)

func IsValid[T any](payload T) error {
	validate := validator.New()

	// Custom validators
	validate.RegisterValidation("password", validation.PasswordValidator)

	err := validate.Struct(payload)
	if err != nil {
		return err
	}
	return nil
}
