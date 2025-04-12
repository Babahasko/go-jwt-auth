package say

import (
	"fmt"
	"net/http"

	"github.com/Babahasko/go-jwt-auth/configs"
	"github.com/Babahasko/go-jwt-auth/pkg/middleware"
	"github.com/Babahasko/go-jwt-auth/pkg/res"
)

type SayHandlerDeps struct{
	Config *configs.Config
}

type SayHandler struct {
}

func NewSayHandler(router *http.ServeMux, deps *SayHandlerDeps) {
	handler := &SayHandler{}
	router.Handle("GET /say/hello", middleware.IsAuthed(handler.Hi(),deps.Config))
	router.HandleFunc("GET /say/bye", handler.Bye())
}

func (handler *SayHandler) Hi() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		contextPayload, _ := r.Context().Value(middleware.ContextEmailKey).(string)
		resultString := fmt.Sprintf("Oh Hi! %v", contextPayload)
		res.Json(w, resultString, http.StatusOK)
	}
}

func (handler *SayHandler) Bye() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		res.Json(w, "Bye bye!", http.StatusOK)
	}
}