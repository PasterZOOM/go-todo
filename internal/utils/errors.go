package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

// ValidationError represents a single validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrorsResponse represents the response for validation errors
type ValidationErrorsResponse struct {
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors"`
}

// HandleValidationErrors processes the validation errors and returns a structured response
func HandleValidationErrors(err error) ValidationErrorsResponse {
	var validationErrors []ValidationError

	var fieldErrors validator.ValidationErrors
	if errors.As(err, &fieldErrors) {
		for _, fieldError := range fieldErrors {
			var detailedMessage string

			switch fieldError.Tag() {
			case "required":
				detailedMessage = fmt.Sprintf("%s is required", fieldError.Field())
			case "email":
				detailedMessage = fmt.Sprintf("%s must be a valid email address", fieldError.Field())
			default:
				detailedMessage = fmt.Sprintf("Invalid value for %s", fieldError.Field())
			}

			validationErrors = append(
				validationErrors, ValidationError{
					Field:   fieldError.Field(),
					Message: detailedMessage,
				},
			)
		}
	}

	return ValidationErrorsResponse{
		Message: "Invalid input data",
		Errors:  validationErrors,
	}
}
