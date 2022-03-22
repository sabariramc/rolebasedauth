package tenant

import "sabariram.com/rolebasedauth/pkg/model"

type CreateClaimDTO struct {
	Claim       string `json:"claim" validate:"nonzero, min=3, max=40, regexp=^[a-z]+(?:.[a-z]+)+$"`
	Description string `json:"description"  validate:"nonzero"`
}

type CreateTenantDTO struct {
	Name               string                           `json:"name" validate:"nonzero, min=3, max=40"`
	BaseURL            string                           `json:"baseURL" validate:"nonzero, url, min=3, max=40"`
	Claims             []*CreateClaimDTO                `json:"claims" validate:"nonzero"`
	AuthenticationType []*model.CreateAuthenticationDTO `json:"authenticationType" validate:"nonzero"`
}
