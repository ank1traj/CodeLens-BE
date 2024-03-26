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

func IsValidIntegerLength(minValue, maxValue, desiredLength int64) error {
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
