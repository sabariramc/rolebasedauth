package middleware

import (
	"net/http"

	"sabariram.com/goserverbase/constant"
	c "sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/errors"
	"sabariram.com/rolebasedauth/pkg/constants"
)

func RequireClaim(claim constants.Claim) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var err error
			reqClaims := r.Context().Value(c.ClaimsKey)
			claimList, ok := reqClaims.([]constants.Claim)
			if ok {
				for _, v := range claimList {
					if claim == v {
						f(w, r)
						return
					}
				}
			}
			err = errors.NewCustomError("NOT_AUTHORIZED", "User is not authorized", nil)
			w.Header().Set(constant.HeaderContentType, constant.ContentTypeJSON)
			w.WriteHeader(http.StatusForbidden)
			b := err.Error()
			_, err = w.Write([]byte(b))
			if err != nil {
				panic(err)
			}

		}
	}
}
