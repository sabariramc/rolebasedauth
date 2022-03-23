package model

import "sabariram.com/rolebasedauth/pkg/constants"

type AuthConfiguration map[string]interface{}

type AllowedAuthentication struct {
	Type          constants.AuthenticationType `json:"type" bson:"type"`
	Configuration AuthConfiguration            `json:"configuration" bson:"configuration"`
}

type Authentication struct {
	AllowedAuthentication
	UserLoginId string `json:"userLoginId" bson:"userLoginId"`
}

type AuthenticationDTO struct {
	Type          constants.AuthenticationType `json:"type" validate:"nonzero, authtype"`
	Configuration AuthConfiguration            `json:"configuration" validate:"nonzero"`
}
