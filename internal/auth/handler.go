package auth

import (
	"net/http"

	"github.com/Babahasko/go-jwt-auth/configs"
	"github.com/Babahasko/go-jwt-auth/pkg/jwt_helper"
	"github.com/Babahasko/go-jwt-auth/pkg/req"
	"github.com/Babahasko/go-jwt-auth/pkg/res"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		email, err := handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		j := jwt_helper.NewJWT(handler.Config.Auth.PrivateKeyFile, handler.Config.Auth.PublicKeyFile)

		token, err := j.Create(jwt_helper.JWTData{Email: email,})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		loginResponseBody := &LoginResponse{
			Token: token,
		}
		res.Json(w, loginResponseBody, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		j := jwt_helper.NewJWT(handler.Config.Auth.PrivateKeyFile, handler.Config.Auth.PublicKeyFile)
		token, err := j.Create(jwt_helper.JWTData{Email: email,})
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody := &RegisterResponse{
			Token: token,
		}

		res.Json(w, responseBody, http.StatusCreated)
	}
}
