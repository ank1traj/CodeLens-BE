package utils

import "errors"

func ValidateStringInput(length int64, count int64) error {
	if length <= 0 {
		return errors.New("desired length should be greater than zero")
	}
	if count <= 0 {
		return errors.New("desired count should be greater than zero")
	}
	return nil
}
