package tenant

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/goserverbase/utils"
	"sabariram.com/rolebasedauth/pkg/model"
	"sabariram.com/rolebasedauth/pkg/utility"
)

type Claim struct {
	mongo.BaseMongoDocument `bson:"inline"`
	ClaimId                 string `json:"claimId" bson:"claimId"`
	Claim                   string `json:"claim" bson:"claim"`
	Description             string `json:"description" bson:"description"`
}

type Tenant struct {
	mongo.BaseMongoModel `bson:",inline"`
	TenantId             string                         `json:"tenantId" bson:"tenantId"`
	Name                 string                         `json:"name" bson:"name"`
	BaseURL              string                         `json:"baseURL" bson:"baseURL"`
	Claims               []*Claim                       `json:"claims" bson:"claims"`
	AuthenticationType   []*model.AllowedAuthentication `json:"authenticationType" bson:"authenticationType"`
	IsActive             bool                           `json:"isActive" bson:"isActive"`
}

func (t *Tenant) Create(ctx context.Context, db *mongo.Mongo) error {
	coll := db.NewCollection("Tenant")
	actor := ctx.Value(constant.ActorIdKey).(string)
	actionAt := time.Now()
	t.CreatedAt = actionAt
	t.UpdatedAt = actionAt
	t.CreatedBy = actor
	t.UpdatedBy = actor
	t.TenantId = utils.GetRandomString(14, "tenant")
	t.IsActive = true
	for _, v := range t.Claims {
		v.CreatedAt = actionAt
		v.UpdatedAt = actionAt
		v.CreatedBy = actor
		v.UpdatedBy = actor
		v.ClaimId = utils.GetRandomString(15, "claim")
	}
	res, err := coll.InsertOne(ctx, t)
	if err != nil {
		return fmt.Errorf("Tenant.Create: %w", err)
	}
	t.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (t *Tenant) Update(ctx context.Context, db *mongo.Mongo, doc *UpdateTenantDTO) error {
	actor := ctx.Value(constant.ActorIdKey).(string)
	actionAt := time.Now()
	t.UpdatedAt = actionAt
	t.UpdatedBy = actor
	u := make(map[string]any)
	if len(doc.Claims) > 0 {
		claims := make([]*Claim, len(doc.Claims))
		err := utility.JsonTransformer(doc.Claims, claims)
		if err != nil {
			return fmt.Errorf("Tenant.Update: %w", err)
		}
		for _, v := range claims {
			if v.ClaimId == "" {
				v.ClaimId = utils.GetRandomString(15, "claim")
				v.CreatedAt = actionAt
				v.CreatedBy = actor
			}
			v.UpdatedAt = actionAt
			v.UpdatedBy = actor
		}
		u["claims"] = claims
	}
	if len(doc.AuthenticationType) > 0 {
		u["authenticationType"] = doc.AuthenticationType
	}
	if len(u) == 0 {
		return nil
	}
	coll := db.NewCollection("Tenant")
	_, err := coll.UpdateOne(ctx, map[string]string{"tenantId": t.TenantId}, map[string]any{"$set": u})
	if err != nil {
		return fmt.Errorf("Tenant.Update: %w", err)
	}
	return nil
}

func (t *Tenant) Delete(ctx context.Context, db *mongo.Mongo) error {
	coll := db.NewCollection("Tenant")
	actor := ctx.Value(constant.ActorIdKey).(string)
	t.UpdatedAt = time.Now()
	t.UpdatedBy = actor
	_, err := coll.UpdateOne(ctx, map[string]string{"tenantId": t.TenantId}, map[string]any{"$set": map[string]bool{"isActive": false}})
	if err != nil {
		return fmt.Errorf("Tenant.Delete: %w", err)
	}
	return nil
}

func (t *Tenant) Get(ctx context.Context, db *mongo.Mongo) error {
	coll := db.NewCollection("Tenant")
	cur := coll.FindOne(ctx, map[string]string{"tenantId": t.TenantId})
	return cur.Decode(t)
}

func List(ctx context.Context, db *mongo.Mongo, filter interface{}, opts ...*options.FindOptions) ([]*Tenant, error) {
	coll := db.NewCollection("Tenant")
	ires, err := coll.FindFetch(ctx, createContainer, filter, opts...)
	var res []*Tenant
	if err != nil {
		return nil, fmt.Errorf("Tenant.List: %w", err)
	}
	res = make([]*Tenant, len(ires))
	for i, v := range ires {
		res[i] = v.(*Tenant)
	}
	return res, nil
}

func createContainer(n int) []interface{} {
	res := make([]interface{}, n)
	for i := 0; i < n; i++ {
		res[i] = &Tenant{}
	}
	return res
}
