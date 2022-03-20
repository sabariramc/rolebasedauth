package model

import "sabariram.com/goserverbase/db/mongo"

type Role struct {
	mongo.BaseMongoModel
	TenantId         string   `json:"tenantId" bson:"tenantId"`
	RoleId           string   `json:"roleId" bson:"roleId"`
	RoleName         string   `json:"roleName" bson:"roleName"`
	IsActive         bool     `json:"isActive" bson:"isActive"`
	Claims           []string `json:"claims" bson:"claims"`
	MaxParallelLogin int      `json:"maxParallelLogin" bson:"maxParallelLogin"`
}
