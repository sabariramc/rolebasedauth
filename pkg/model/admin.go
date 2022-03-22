package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID       primitive.ObjectID `json:"-" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	ApiKey   string             `json:"apiKey" bson:"apiKey"`
	IsActive bool               `json:"isActive" bson:"isActive"`
	Claims   []string           `json:"claims" bson:"claims"`
	TenantId string             `json:"tenantId" bson:"tenantId"`
}
