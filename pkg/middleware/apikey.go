package middleware

import (
	"context"
	"net/http"

	"sabariram.com/goserverbase/constant"
	"sabariram.com/goserverbase/db/mongo"
	"sabariram.com/goserverbase/errors"
	"sabariram.com/rolebasedauth/pkg/model"
)

func RequireClaim(admin *mongo.Collection) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			
		}
	}
}
