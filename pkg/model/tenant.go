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
	TenantId           string          `json:"tenantId" bson:"tenantId"`
	Name               string          `json:"name" bson:"name"`
	Path               string          `json:"path" bson:"path"`
	BaseURL            string          `json:"baseURL" bson:"baseURL"`
	Claims             []*Claim        `json:"claims" bson:"claims"`
	AuthenticationType *Authentication `json:"authenticationType" bson:"authenticationType"`
}
