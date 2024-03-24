package utils

import "errors"

func DesireCountLimit(count int64) error {
	if count > 100000 {
		return errors.New("for now we support only upto 10^5")
	}
	return nil
}
