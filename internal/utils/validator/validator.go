package validator

import (
	"github.com/go-playground/validator/v10"
)

// Validator instance
var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
}

// ValidateStruct validates a struct based on the tags
func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}
