package model

import "sabariram.com/rolebasedauth/pkg/constants"

type AllowedAuthentication struct {
	Type          constants.AuthenticationType `json:"type" bson:"type"`
	Configuration map[string]string            `json:"configuration" bson:"configuration"`
}

type Authentication struct {
	AllowedAuthentication
	TenantUserIdentifier string `json:"tenantUserIdentifier" bson:"tenantUserIdentifier"`
}
