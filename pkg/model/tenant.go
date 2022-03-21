package model

import "sabariram.com/goserverbase/db/mongo"

type Claim struct {
	mongo.BaseMongoDocument
	ClaimId     string `json:"claimId" bson:"claimId"`
	Claim       string `json:"claim" bson:"claim"`
	Description string `json:"description" bson:"description"`
}

type Tenant struct {
	mongo.BaseMongoModel
	TenantId           string                   `json:"tenantId" bson:"tenantId"`
	Name               string                   `json:"name" bson:"name"`
	Path               string                   `json:"path" bson:"path"`
	BaseURL            string                   `json:"baseURL" bson:"baseURL"`
	Claims             []*Claim                 `json:"claims" bson:"claims"`
	AuthenticationType []*AllowedAuthentication `json:"authenticationType" bson:"authenticationType"`
}

type CreateClaimDTO struct {
	Claim       string `json:"claim" validate:"nonzero, min=3, max=40, regexp=^[a-z]+(?:.[a-z]+)+$"`
	Description string `json:"description"  validate:"nonzero"`
}

type CreateTenantDTO struct {
	Name               string                     `json:"name" validate:"nonzero, min=3, max=40"`
	Path               string                     `json:"path" validate:"nonzero, min=3, max=40, regexp=^[a-zA-Z]*$"`
	BaseURL            string                     `json:"baseURL" validate:"nonzero, min=3, max=40, regexp=^(http|https)://[a-z0-9-]+(?:.[a-z0-9-]+)+.([a-z]{2}|[a-z]{3})$"`
	Claims             []*CreateClaimDTO          `json:"claims" validate:"nonzero"`
	AuthenticationType []*CreateAuthenticationDTO `json:"authenticationType" validate:"nonzero"`
}
