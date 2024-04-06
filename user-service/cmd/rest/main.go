package main

import (
	"fmt"
	"go-microservice/user-service/api"
	"go-microservice/user-service/internal/config"
	"go-microservice/user-service/internal/database"
	"net/http"
)

func main() {
	router := api.NewRouter()

	config.LoadConfig()

	database.Connect()

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	fmt.Println("Starting up server...")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Failed to start the server")

		return
	}
	fmt.Println("Server Stopped")
}
