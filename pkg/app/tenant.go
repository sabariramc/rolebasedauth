package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sabariram.com/goserverbase/baseapp"
	"sabariram.com/goserverbase/errors"
	"sabariram.com/goserverbase/utils"
	"sabariram.com/rolebasedauth/pkg/constants"
	"sabariram.com/rolebasedauth/pkg/model/admin"
	"sabariram.com/rolebasedauth/pkg/model/tenant"
)

func (rbac *RoleBasedAuthentication) CreateTenant() http.HandlerFunc {
	var body tenant.CreateTenantDTO
	return rbac.b.JSONResponder(&body, func(r *http.Request) (statusCode int, res interface{}, err error) {
		if err := rbac.validator.Validate(body); err != nil {
			return http.StatusBadRequest, nil, errors.NewCustomError("INVALID_PAYLOAD", "Error in payload", err)
		}
		t := &tenant.Tenant{}
		a := &admin.Admin{}
		err = utils.JsonTransformer(body, t)
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

func (rbac *RoleBasedAuthentication) SearchTenant() http.HandlerFunc {
	return rbac.b.JSONResponder(nil, func(r *http.Request) (statusCode int, res interface{}, err error) {
		tenantId := r.Context().Value(constants.ContextVariableTenantIdKey)
		if tenantId != nil {
			tList, err := tenant.List(r.Context(), rbac.db, map[string]string{"tenantId": tenantId.(string)})
			if err != nil {
				panic(fmt.Errorf("RoleBasedAuthentication.SearchTenant : %w", err))
			}
			if len(tList) == 0 {
				return http.StatusBadRequest, nil, nil
			}
			return http.StatusOK, tList[0], nil
		}
		return http.StatusBadRequest, nil, nil
	})
}

func (rbac *RoleBasedAuthentication) ListTenant() http.HandlerFunc {
	return rbac.b.JSONResponder(nil, func(r *http.Request) (statusCode int, res interface{}, err error) {
		var qp baseapp.Filter
		err = schema.NewDecoder().Decode(&qp, r.URL.Query())
		if err != nil {
			return http.StatusBadRequest, nil, fmt.Errorf("RoleBasedAuthentication.ListTenant : %w", err)
		}
		err = baseapp.SetDefaultPagination(&qp, "name")
		if err != nil {
			return http.StatusBadRequest, nil, fmt.Errorf("RoleBasedAuthentication.ListTenant : %w", err)
		}
		sortOrder := -1
		if *qp.Asc {
			sortOrder = 1
		}
		offset := (qp.PageNo - 1) * qp.Limit
		page := options.Find().SetLimit(qp.Limit).SetSkip(offset).SetSort(bson.D{{qp.SortBy, sortOrder}})
		tList, err := tenant.List(r.Context(), rbac.db, nil, page)
		if err != nil {
			panic(fmt.Errorf("RoleBasedAuthentication.ListTenant : %w", err))
		}
		return http.StatusOK, tList, nil
	})
}

func (rbac *RoleBasedAuthentication) GetTenant() http.HandlerFunc {
	return rbac.b.JSONResponder(nil, func(r *http.Request) (statusCode int, res interface{}, err error) {
		t := &tenant.Tenant{}
		t.TenantId = mux.Vars(r)["tenantId"]
		err = t.Get(r.Context(), rbac.db)
		if err != nil {
			panic(fmt.Errorf("RoleBasedAuthentication.GetTenant : %w", err))
		}
		return http.StatusOK, t, nil
	})
}

func (rbac *RoleBasedAuthentication) UpdateTenant() http.HandlerFunc {
	var body tenant.UpdateTenantDTO
	return rbac.b.JSONResponder(&body, func(r *http.Request) (statusCode int, res interface{}, err error) {
		return 100, nil, nil
	})
}
