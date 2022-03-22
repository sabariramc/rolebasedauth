package middleware

import (
	"net/http"

	"github.com/gorilla/mux"
	"sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/errors"
	"sabariram.com/rolebasedauth/pkg/constants"
)

func RequireTenant(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		params := mux.Vars(r)
		value, ok := r.Context().Value(constants.TenantPathKey).(string)
		if ok && value != "" && params["tenantId"] == value {
			f(w, r)
		}
		err = errors.NewCustomError("NOT_AUTHORIZED", "User is not authorized for the tenant", nil)
		w.Header().Set(constant.HeaderContentType, constant.ContentTypeJSON)
		w.WriteHeader(http.StatusForbidden)
		b := err.Error()
		_, err = w.Write([]byte(b))
		if err != nil {
			panic(err)
		}
	}
}
