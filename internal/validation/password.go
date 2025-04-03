package validation

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func PasswordValidator(fl validator.FieldLevel) bool {
	min_length := 8
	max_length := 20
	pass :=fl.Field().String()
	if len(pass) < min_length {
		return false
	}
	if len(pass) > max_length {
		return false
	}
	hasUpperCase := regexp.MustCompile(`.*[A-Z].*`).MatchString(pass)

	hasLowerCase := regexp.MustCompile(`.*[a-z].*`).MatchString(pass)

	hasNumber := regexp.MustCompile(`.*[0-9].*`).MatchString(pass)

	if !hasNumber || !hasLowerCase || !hasUpperCase{
		return false
	}
	return true
}