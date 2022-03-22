package app

import (
	"regexp"

	"gopkg.in/validator.v2"
	"sabariram.com/goserverbase/errors"
	"sabariram.com/rolebasedauth/pkg/constants"
)

func (r *RoleBasedAuthentication) registerValidator() {
	r.validator.SetValidationFunc("authtype", ValidateEnum(constants.AuthTypeList))
	r.validator.SetValidationFunc("url", ValidateURL)
}

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
		return errors.NewCustomError("BAD_PARAM", "Invalid value for param", map[string]interface{}{
			"expectedValues": enumList,
		})
	}
}

func ValidateURL(v interface{}, param string) error {
	val, ok := v.(string)
	if !ok {
		return validator.ErrUnsupported
	}
	match, err := regexp.Match("^(http|https)://[a-z0-9-]+(?:.[a-z0-9-]+)+\\.([a-z]{2,3})$", []byte(val))
	if err == nil && match {
		return nil
	}
	return errors.NewCustomError("BAD_PARAM", "Invalid value for param", nil)
}
