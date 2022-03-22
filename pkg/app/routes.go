package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"sabariram.com/rolebasedauth/pkg/constants"
	"sabariram.com/rolebasedauth/pkg/middleware"
)

func (r *RoleBasedAuthentication) registerBookStoreRoutes(router *mux.Router) {
	router = router.PathPrefix("/tenant").Subrouter()
	router.Methods(http.MethodPost).HandlerFunc(middleware.RequireClaim(constants.TenantCreate)(r.CreateTenant()))

	router = router.PathPrefix("/{tenantId}").Subrouter()

	// u := router.PathPrefix("/user").Subrouter()

}
