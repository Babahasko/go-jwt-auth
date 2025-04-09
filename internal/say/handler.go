package say

import (
	"net/http"

	"github.com/Babahasko/go-jwt-auth/pkg/middleware"
	"github.com/Babahasko/go-jwt-auth/pkg/res"
)

type SayHandler struct {
}

func NewSayHandler(router *http.ServeMux) {
	handler := &SayHandler{

	}
	router.Handle("GET /say/hello", middleware.IsAuthed(handler.Hi()))
	router.HandleFunc("GET /say/bye", handler.Bye())
}

func (handler *SayHandler) Hi() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, "Oh. Hi!", http.StatusOK)
	}
}

func (handler *SayHandler) Bye() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, "Bye bye!", http.StatusOK)
	}
}