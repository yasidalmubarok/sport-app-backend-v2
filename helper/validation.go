package helper

import (
	"github.com/go-playground/validator/v10"
)

// Buat instance validator
var validate = validator.New()

func ValidateStruct(s interface{}) Error {
	// Validasi struct
	err := validate.Struct(s)

	if err != nil {
		// Ambil pesan error pertama
		for _, e := range err.(validator.ValidationErrors) {
			return NewBadRequestError(e.Error())
		}
	}

	return nil
}
