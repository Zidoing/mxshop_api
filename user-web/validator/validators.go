package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateMobile(field validator.FieldLevel) bool {
	mobile := field.Field().String()

	ok, _ := regexp.MatchString(`^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
