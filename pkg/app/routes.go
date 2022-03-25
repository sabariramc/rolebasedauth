package app

import (
	"net/http"

	"sabariram.com/goserverbase/baseapp"
	"sabariram.com/rolebasedauth/pkg/constants"
	"sabariram.com/rolebasedauth/pkg/middleware"
)

func (rbac *RoleBasedAuthentication) Routes() *baseapp.APIRoute {
	return &baseapp.APIRoute{
		"/tenant": &baseapp.APIResource{
			Handlers: map[string]*baseapp.APIHandler{
				http.MethodPost: {
					Func: middleware.RequireClaim(constants.TenantCreate)(rbac.CreateTenant()),
				},
				http.MethodGet: {
					Func: middleware.RequireClaim(constants.TenantList)(rbac.ListTenant()),
				},
			},
			SubResource: map[string]*baseapp.APIResource{
				"/search": {
					Handlers: map[string]*baseapp.APIHandler{
						http.MethodGet: {
							Func: middleware.RequireClaim(constants.TenantSearch)(rbac.SearchTenant()),
						},
					},
				},
				"/{tenantId}": {
					Handlers: map[string]*baseapp.APIHandler{
						http.MethodGet: {
							Func: middleware.RequireClaim(constants.TenantGet)(middleware.RequireTenant(rbac.GetTenant())),
						},
					},
				},
			},
		},
	}
}
