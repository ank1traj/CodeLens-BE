package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

func DesireCountLimit(count int64) error {
	if count > 100000 {
		return errors.New("for now we support only upto 10^5")
	}
	return nil
}

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

func ValidateStruct(s interface{}, validate *validator.Validate) error {
	err := validate.Struct(s)
	if err != nil {
		var validationErrors []ValidationError
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Field: err.Field(),
				Error: formatErrorMessage(err),
			})
		}
		return fmt.Errorf("validation errors: %s", validationErrorsToString(validationErrors))
	}
	return nil
}

func formatErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("field %s is required", err.Field())
	case "min":
		return fmt.Sprintf("field %s must be at least %s", err.Field(), err.Param())
	default:
		return fmt.Sprintf("validation failed on field %s", err.Field())
	}
}

func validationErrorsToString(errors []ValidationError) string {
	var errorStrings []string
	for _, err := range errors {
		errorStrings = append(errorStrings, fmt.Sprintf("%s: %s", err.Field, err.Error))
	}
	return strings.Join(errorStrings, ", ")
}
