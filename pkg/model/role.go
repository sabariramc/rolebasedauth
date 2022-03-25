package model

import (
	"context"
	"fmt"
	"time"

	"sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/goserverbase/utils"
)

type RoleDTO struct {
	RoleName         *string  `json:"roleName,omitempty" validate:"nonzero, min=3, max=40"`
	Claims           []string `json:"claims,omitempty"  validate:"nonzero"`
	MaxParallelLogin *int     `json:"maxParallelLogin,omitempty" validate:"nonzero, min=1, max=10"`
}

type Role struct {
	mongo.BaseMongoModel `bson:",inline"`
	TenantId             string   `json:"tenantId" bson:"tenantId"`
	RoleId               string   `json:"roleId" bson:"roleId"`
	Name                 *string  `json:"name" bson:"nsame,omitempty"`
	IsActive             bool     `json:"isActive" bson:"isActive"`
	Claims               []string `json:"claims" bson:"claims,omitempty"`
	MaxParallelLogin     *int     `json:"maxParallelLogin" bson:"maxParallelLogin,omitempty"`
}

func (r *Role) Create(ctx context.Context, db *mongo.Mongo) error {
	coll := db.NewCollection("Role")
	actor := ctx.Value(constant.ActorIdKey).(string)
	actionAt := time.Now()
	r.CreatedAt = actionAt
	r.UpdatedAt = actionAt
	r.CreatedBy = actor
	r.UpdatedBy = actor
	r.RoleId = utils.GetRandomString(14, "role")
	r.IsActive = true
	_, err := coll.InsertOne(ctx, r)
	if err != nil {
		return fmt.Errorf("Role.Create: %w", err)
	}
	return nil
}
