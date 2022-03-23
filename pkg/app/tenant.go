package app

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"sabariram.com/goserverbase/errors"
	"sabariram.com/rolebasedauth/pkg/model/admin"
	"sabariram.com/rolebasedauth/pkg/model/tenant"
	"sabariram.com/rolebasedauth/pkg/utility"
)

func (rbac *RoleBasedAuthentication) CreateTenant() http.HandlerFunc {
	var body tenant.CreateTenantDTO
	return rbac.b.JSONResponder(&body, func(r *http.Request) (statusCode int, res interface{}, err error) {
		if err := rbac.validator.Validate(body); err != nil {
			return http.StatusBadRequest, nil, errors.NewCustomError("INVALID_PAYLOAD", "Error in payload", err)
		}
		t := &tenant.Tenant{}
		a := &admin.Admin{}
		err = utility.JsonTransformer(body, t)
		if err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("RoleBasedAuthentication.CreateTenant Transformation: %w", err)
		}
		session, err := rbac.db.GetClient().StartSession()
		if err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("RoleBasedAuthentication.CreateTenant StartSession : %w", err)
		}
		if err = session.StartTransaction(); err != nil {
			return http.StatusInternalServerError, nil, fmt.Errorf("RoleBasedAuthentication.CreateTenant StartTransaction : %w", err)
		}
		ctx := r.Context()
		var apiKey string
		err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
			err = t.Create(sc, rbac.db)
			if err != nil {
				session.AbortTransaction(sc)
				return fmt.Errorf("RoleBasedAuthentication.CreateTenant NewTenant : %w", err)
			}
			tenantId := t.TenantId

			apiKey, err = a.Create(sc, rbac.db, tenantId, t.Name)
			if err != nil {
				session.AbortTransaction(sc)
				return fmt.Errorf("RoleBasedAuthentication.CreateTenant NewAdmin : %w", err)
			}
			session.CommitTransaction(sc)
			return nil
		})
		session.EndSession(ctx)
		if err != nil {
			if mongo.IsDuplicateKeyError(err) {
				return http.StatusBadRequest, nil, errors.NewCustomError("DUPLICATE_DATA", "Error in payload", err)
			}
			rbac.log.Error(ctx, "Error creating Tenant", err)
			panic(err)
		}
		return http.StatusOK, map[string]any{"tenant": t, "admin": a, "adminKey": apiKey}, nil
	})
}
