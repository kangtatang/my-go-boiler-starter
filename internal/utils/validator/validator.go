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

// package validator

// import "github.com/go-playground/validator/v10"

// type Validator struct {
// 	validate *validator.Validate
// }

// // NewValidator membuat instance Validator baru
// func NewValidator() *Validator {
// 	return &Validator{validate: validator.New()}
// }

// // Validate menjalankan validasi pada struct yang diberikan
// func (v *Validator) Validate(i interface{}) error {
// 	return v.validate.Struct(i)
// }
