package utils

import "github.com/go-playground/validator/v10"

func FormatValidationErrors(validationErrors validator.ValidationErrors) map[string]string {
	errorMessages := make(map[string]string)
	for _, err := range validationErrors {
		switch err.Field() {
		case "Email":
			errorMessages["message"] = "email was invalid"
		case "Password":
			errorMessages["message"] = "password is required"
		default:
			errorMessages["message"] = "invalid input"
		}
	}
	return errorMessages
}
