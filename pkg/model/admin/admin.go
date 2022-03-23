package admin

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/goserverbase/utils"
	"sabariram.com/rolebasedauth/pkg/constants"
)

type Admin struct {
	mongo.BaseMongoModel `bson:",inline"`
	Name                 string            `json:"name" bson:"name"`
	ApiKey               string            `json:"apiKey" bson:"apiKey"`
	IsActive             bool              `json:"isActive" bson:"isActive"`
	Claims               []constants.Claim `json:"claims" bson:"claims"`
	TenantId             string            `json:"tenantId" bson:"tenantId"`
}

func (a *Admin) Create(ctx context.Context, db *mongo.Mongo, tenantId, tenantName string) (string, error) {
	coll := db.NewCollection("Admin")
	actor := ctx.Value(constant.ActorIdKey).(string)
	actionAt := time.Now()
	a.CreatedAt = actionAt
	a.UpdatedAt = actionAt
	a.CreatedBy = actor
	a.UpdatedBy = actor
	a.Name = tenantName
	a.IsActive = true
	a.TenantId = tenantId
	a.Claims = constants.RoleTenantAdmin
	apiKey := utils.GetRandomString(20, "")
	a.ApiKey = utils.GetHash(apiKey)
	res, err := coll.InsertOne(ctx, a)
	if err != nil {
		return "", fmt.Errorf("Admin.Create: %w", err)
	}
	a.ID = res.InsertedID.(primitive.ObjectID)
	return apiKey, nil
}
