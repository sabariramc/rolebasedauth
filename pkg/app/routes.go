package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (r *RoleBasedAuthentication) registerBookStoreRoutes(router *mux.Router) {
	router.HandleFunc("/tenant", r.CreateTenant()).Methods(http.MethodPost)
}
