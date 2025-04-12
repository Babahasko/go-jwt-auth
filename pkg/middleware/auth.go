package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Babahasko/go-jwt-auth/configs"
	"github.com/Babahasko/go-jwt-auth/pkg/jwt_helper"
)

type key string

const (
	ContextEmailKey key = "ContextEmailKey"
)

func IsAuthed(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		jwtPayload, err := jwt_helper.NewJWT(conf.Auth.PrivateKeyFile, conf.Auth.PublicKeyFile).Parse(token)
		
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		r.Context()
		ctx := context.WithValue(r.Context(), ContextEmailKey, jwtPayload.Email)
		rWithCotext := r.WithContext(ctx)

		next.ServeHTTP(w, rWithCotext)
	})
}