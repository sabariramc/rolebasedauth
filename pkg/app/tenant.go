package app

import (
	"net/http"

	"sabariram.com/goserverbase/errors"
	"sabariram.com/rolebasedauth/pkg/model"
)

func (r *RoleBasedAuthentication) CreateTenant() http.HandlerFunc {

	var body model.CreateTenantDTO
	return r.b.JSONResponder(&body, func(req *http.Request) (statusCode int, res interface{}, err error) {
		if err := r.validator.Validate(body); err != nil {
			return http.StatusBadRequest, nil, errors.NewCustomError("INVALID_PAYLOAD", "Error in payload", err)
		}
		return http.StatusAccepted, nil, nil
	})
}
