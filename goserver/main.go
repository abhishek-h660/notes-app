package main

import (
	"log"
	"net/http"
	"notes/server/api"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	godotenv.Load()

	log.Println("Notes Server is starting...")

	s := api.NewServer()

	// Enabling cross origin for API testing
	headers := handlers.AllowedHeaders([]string{"*"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, origins, methods)(s)))
}
