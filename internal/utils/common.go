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
		return fmt.Errorf(validationErrorsToString(validationErrors))
	}
	return nil
}

func formatErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("%s field is required", err.Field())
	case "min":
		return fmt.Sprintf("%s field must be at least %s", err.Field(), err.Param())
	case "gte":
		return fmt.Sprintf("%s field must be greater or equal to %s", err.Field(), err.Param())
	default:
		return fmt.Sprintf("validation failed on field %s", err.Field())
	}
}

func validationErrorsToString(errors []ValidationError) string {
	var errorStrings []string
	for _, err := range errors {
		errorStrings = append(errorStrings, err.Error)
	}
	return strings.Join(errorStrings, ", ")
}

func ValidateMinAndMax(minValue, maxValue int64) error {
	if minValue > maxValue {
		return errors.New("minValue should be less than or equal to maxValue")
	}
	return nil
}
