package utils

import (
	"errors"
)

func ValidateArrayInputs(minValue, maxValue, desiredSize, desiredLength, count int64) error {
	if minValue > maxValue {
		return errors.New("minValue should be less than or equal to maxValue")
	}
	if desiredLength <= 0 {
		return errors.New("desired length should be greater than zero")
	}
	if count <= 0 {
		return errors.New("desired count should be greater than zero")
	}
	if desiredSize <= 0 {
		return errors.New("desired size should be greater than zero")
	}
	return nil
}
