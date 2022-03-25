package app

import (
	"fmt"
	"net/http"

	"sabariram.com/goserverbase/errors"
	"sabariram.com/goserverbase/utils"
	"sabariram.com/rolebasedauth/pkg/constants"
	"sabariram.com/rolebasedauth/pkg/model"
	"sabariram.com/rolebasedauth/pkg/model/tenant"
)

func (rbac *RoleBasedAuthentication) CreateRole() http.HandlerFunc {
	var body model.RoleDTO
	return rbac.b.JSONResponder(body, func(r *http.Request) (statusCode int, res interface{}, err error) {
		if err := rbac.validator.Validate(body); err != nil {
			return http.StatusBadRequest, nil, errors.NewCustomError("INVALID_PAYLOAD", "Error in payload", err)
		}
		role := &model.Role{}
		err = utils.JsonTransformer(body, role)
		if err != nil {
			panic(fmt.Errorf("RoleBasedAuthentication.CreateRole Transformation: %w", err))
		}
		role.TenantId = r.Context().Value(constants.ContextVariableTenantIdKey).(string)
		t := &tenant.Tenant{}
		t.TenantId = role.TenantId
		t.Get(r.Context(), rbac.db)
		claimMap := make(map[string]*tenant.Claim, len(t.Claims))
		for _, v := range t.Claims {
			claimMap[v.ClaimId] = v
		}
		invalidClaimId := make([]string, 0)
		ivflag := false
		for _, v := range role.Claims {
			_, ok := claimMap[v]
			if !ok {
				ivflag = true
				invalidClaimId = append(invalidClaimId, v)
			}
		}
		if ivflag {
			return http.StatusBadRequest, nil, errors.NewCustomError("INVALID_CLAIM_ID", "Invalid ClaimId", invalidClaimId)
		}
		role.Create(r.Context(), rbac.db)
		return http.StatusOK, role, nil
	})
}
