package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Babahasko/go-jwt-auth/configs"
	"github.com/Babahasko/go-jwt-auth/internal/auth"
	"github.com/Babahasko/go-jwt-auth/pkg/db"
)

func main() {
	router := http.NewServeMux()
	conf := configs.LoadConfig()

	//Db connection
	_ = db.NewDb(conf)
	
	//Handler
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	fmt.Println("Listen and serve :8081")
	log.Fatal(server.ListenAndServe())
}
