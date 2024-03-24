package utils

import (
	"errors"
	"math"
	"strconv"
)

func CountIntegers(num int64) int64 {
	absNum := int64(math.Abs(float64(num)))
	var numStr = strconv.Itoa(int(absNum))
	return int64(len(numStr))
}

func IsValidLength(minValue, maxValue, desiredLength int64) error {
	minLengthValue := CountIntegers(minValue)
	maxLengthValue := CountIntegers(maxValue)

	if minValue < 0 && maxValue > 0 {
		if desiredLength > minLengthValue && desiredLength > maxLengthValue {
			return errors.New("desired length is not valid for the given range")
		}
	} else if minValue < 0 || maxValue < 0 {
		if desiredLength > minLengthValue || desiredLength < maxLengthValue {
			return errors.New("desired length is not valid for the given range")
		}
	} else {
		if desiredLength < minLengthValue || desiredLength > maxLengthValue {
			return errors.New("desired length is not valid for the given range")
		}
	}
	return nil
}

func DesireCountLimit(count int64) error {
	if count > 100000 {
		return errors.New("for now we support only upto 10^5")
	}
	return nil
}

func ValidateInputs(minValue, maxValue, desiredLength, count int64) error {
	if minValue > maxValue {
		return errors.New("minValue should be less than or equal to maxValue")
	}
	if desiredLength <= 0 {
		return errors.New("desired length should be greater than zero")
	}
	if count <= 0 {
		return errors.New("desired count should be greater than zero")
	}
	return nil
}
