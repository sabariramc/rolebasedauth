package app

import (
	"gopkg.in/validator.v2"
	"sabariram.com/goserverbase/errors"
)

func ValidateEnum[K comparable](enumList []K) validator.ValidationFunc {
	return func(v interface{}, param string) error {
		val, ok := v.(K)
		if !ok {
			return validator.ErrUnsupported
		}
		for _, l := range enumList {
			if val == l {
				return nil
			}
		}
		return errors.NewCustomError("BAD_PARAM", "Invalid valie for param", map[string]interface{}{
			"expectedValues": enumList,
		})
	}
}
