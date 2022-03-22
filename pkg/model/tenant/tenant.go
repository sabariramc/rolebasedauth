package tenant

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/rolebasedauth/pkg/model"
)

type Claim struct {
	mongo.BaseMongoDocument
	ClaimId     string `json:"claimId" bson:"claimId"`
	Claim       string `json:"claim" bson:"claim"`
	Description string `json:"description" bson:"description"`
}

type Tenant struct {
	mongo.BaseMongoModel
	TenantId           string                         `json:"tenantId" bson:"tenantId"`
	Name               string                         `json:"name" bson:"name"`
	BaseURL            string                         `json:"baseURL" bson:"baseURL"`
	Claims             []*Claim                       `json:"claims" bson:"claims"`
	AuthenticationType []*model.AllowedAuthentication `json:"authenticationType" bson:"authenticationType"`
}

func Create(ctx context.Context, db *mongo.Mongo, t *Tenant) error {
	coll := db.NewCollection("Tenant")
	actor := ctx.Value(constant.ActorIdKey).(string)
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	t.CreatedBy = actor
	t.UpdatedBy = actor

	res, err := coll.InsertOne(ctx, t)
	if err != nil {
		return err
	}
	t.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
