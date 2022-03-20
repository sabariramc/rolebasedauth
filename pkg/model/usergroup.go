package model

import "sabariram.com/goserverbase/db/mongo"

type UserGroup struct {
	mongo.BaseMongoModel
	TenantId  string `json:"tenantId" bson:"tenantId"`
	GroupName string `json:"groupName" bson:"groupName"`
	GroupId   string `json:"groupId" bson:"groupId"`
	UserId    string `json:"userId" bson:"userId"`
	RoleId    string `json:"roleId" bson:"roleId"`
}
