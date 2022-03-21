package app

import "sabariram.com/rolebasedauth/pkg/constants"

func (r *RoleBasedAuthentication) registerValidator() {
	r.validator.SetValidationFunc("authtype", ValidateEnum(constants.AuthTypeList))
}
