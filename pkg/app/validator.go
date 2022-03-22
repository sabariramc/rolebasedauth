package app

import "sabariram.com/rolebasedauth/pkg/constants"

func (r *RoleBasedAuthentication) registerValidator() {
	r.validator.SetValidationFunc("authtype", ValidateEnum(constants.AuthTypeList))
	r.validator.SetValidationFunc("url", ValidateURL)
}
