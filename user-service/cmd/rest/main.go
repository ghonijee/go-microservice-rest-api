package main

import (
	"fmt"
	"go-microservice/user-service/api"
	"go-microservice/user-service/internal/common"
	"go-microservice/user-service/internal/config"
	"go-microservice/user-service/internal/database"
	"net/http"
)

func main() {
	config.LoadConfig()
	common.LoadValidator()
	database.Connect()

	router := api.NewRouter()

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	fmt.Println("Starting up server...")
	fmt.Println("Listening on port 8000")

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Failed to start the server")
	}
	defer database.DB.Close()
	fmt.Println("Server Stopped")
}
