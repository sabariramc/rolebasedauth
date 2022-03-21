package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"sabariram.com/rolebasedauth/pkg/middleware"
)

func (r *RoleBasedAuthentication) registerBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/tenant", middleware.RequireApiKey(r.adminAuth)(r.CreateTenant())).Methods(http.MethodPost)
}
