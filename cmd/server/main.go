package main

import (
	"log"
	"net/http"
	"os"

	"../../configs"
	"../../routes"
)

func main() {
	config := configs.LoadConfig(os.Getenv("APP_ENV"))
	router := routes.NewRouter()

	log.Printf("Server is listening on %s", config.Port)
	if err := http.ListenAndServe(config.Port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}