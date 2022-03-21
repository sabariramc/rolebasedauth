package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	ApiKey   string             `bson:"apiKey"`
	IsActive bool               `bson:"isActive"`
	Claims   []string           `bson:"claims"`
}
