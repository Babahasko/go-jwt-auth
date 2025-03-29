package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Babahasko/go-jwt-auth/internal/auth"
)

func main() {
	router := http.NewServeMux()
	auth.NewAuthHandler(router)

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	fmt.Println("Listen and serve :8081")
	log.Fatal(server.ListenAndServe())
}
