package middleware

import (
	"context"
	"net/http"

	"sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/goserverbase/errors"
	"sabariram.com/rolebasedauth/pkg/model"
)

func RequireApiKey(admin *mongo.Collection) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("x-api-key")
			var err error
			if apiKey != "" {
				cur := admin.FindOne(r.Context(), map[string]interface{}{"apiKey": apiKey, "isActive": true})
				val := &model.Admin{}
				err = cur.Decode(val)
				if err == nil {
					r = r.WithContext(context.WithValue(r.Context(), constant.ActorIdKey, val.Name))
					f(w, r)
					return
				}
				err = errors.NewCustomError("INVALID_API_KEY", "invalid api key", nil)
			} else {
				err = errors.NewCustomError("MISSING_API_KEY", "requires api key", nil)
			}
			w.Header().Set(constant.HeaderContentType, constant.ContentTypeJSON)
			w.WriteHeader(http.StatusUnauthorized)
			b := err.Error()
			_, err = w.Write([]byte(b))
			if err != nil {
				panic(err)
			}
		}
	}
}