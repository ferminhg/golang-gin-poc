package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateCoolName(fl validator.FieldLevel) bool {
	return strings.Contains(fl.Field().String(), "Muka")
}
