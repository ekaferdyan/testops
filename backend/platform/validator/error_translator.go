package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// TranslateError menerjemahkan error teknis menjadi []ErrorResponse yang rapi
func TranslateError(err error) []ErrorResponse {
	var errors []ErrorResponse

	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range validationErrors {
			errors = append(errors, ErrorResponse{
				Field:   strings.ToLower(fe.Field()),
				Message: getErrorMessage(fe),
			})
		}
	}

	return errors
}

// getErrorMessage adalah "kamus" penerjemah kita
func getErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Wajib diisi"
	case "username":
		return "Username tidak valid"
	case "min":
		return fmt.Sprintf("Minimal harus %s karakter", fe.Param())
	case "max":
		return fmt.Sprintf("Maksimal %s karakter", fe.Param())
	default:
		return "Input tidak valid"
	}
}
