package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Babahasko/go-jwt-auth/configs"
	"github.com/Babahasko/go-jwt-auth/pkg/jwt"
)

func IsAuthed(next http.Handler, conf *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		payload, err := jwt.NewJWT(conf.Auth.PrivateKeyFile, conf.Auth.PublicKeyFile).Parse(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Println(payload)
		next.ServeHTTP(w, r)
	})
}