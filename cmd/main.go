package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Babahasko/go-jwt-auth/configs"
	"github.com/Babahasko/go-jwt-auth/internal/auth"
	"github.com/Babahasko/go-jwt-auth/internal/say"
	"github.com/Babahasko/go-jwt-auth/internal/user"
	"github.com/Babahasko/go-jwt-auth/pkg/db"
	"github.com/Babahasko/go-jwt-auth/pkg/middleware"
)

func main() {
	router := http.NewServeMux()
	conf := configs.LoadConfig()

	//Db connection
	db := db.NewDb(conf)

	//Repositories
	userRepository := user.NewUserRepository(db)

	//Service
	authService := auth.NewAuthService(userRepository)
	
	//Handler
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: conf,
		AuthService: authService,
	})

	say.NewSayHandler(router, &say.SayHandlerDeps{
		Config: conf,
	})
	
	// stack Middleware
	stack := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	server := http.Server{
		Addr: ":8081",
		Handler: stack(router),
	}
	fmt.Println("Listen and serve :8081")
	log.Fatal(server.ListenAndServe())
}
