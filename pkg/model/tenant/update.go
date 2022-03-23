package tenant

import "sabariram.com/rolebasedauth/pkg/model"

type UpdateClaimDTO struct {
	ClaimId     string `json:"claimId,omitempty" bson:"claimId,omitempty"`
	Claim       string `json:"claim,omitempty" bson:"claim,omitempty" validate:"nonzero, min=3, max=40, regexp=^[a-z]+(?:.[a-z]+)+$"`
	Description string `json:"description,omitempty" bson:"description,omitempty"  validate:"nonzero"`
}

type UpdateTenantDTO struct {
	Claims             []*UpdateClaimDTO          `json:"claims,omitempty" bson:"claims,omitempty"`
	AuthenticationType []*model.AuthenticationDTO `json:"authenticationType,omitempty" bson:"authenticationType, omitempty"`
}
