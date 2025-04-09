package auth

import (
	"net/http"

	"github.com/Babahasko/go-jwt-auth/configs"
	"github.com/Babahasko/go-jwt-auth/pkg/req"
	"github.com/Babahasko/go-jwt-auth/pkg/res"
)

type AuthHandlerDeps struct{
	*configs.Config
}

type AuthHandler struct{
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[LoginRequest](w,r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := LoginResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusOK)
	}	
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := req.HandleBody[RegisterRequest](w,r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		data := RegisterResponse{
			Token: "123",
		}
		res.Json(w, data, http.StatusCreated)
	}
}