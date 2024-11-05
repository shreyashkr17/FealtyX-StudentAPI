package main

import (
	"log"
	"net/http"
	"os"
	"student-api/api/handlers"
	"student-api/api/middleware"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/students/", handlers.StudentHandler)

	loggedMux := middleware.LoggingMiddleware(mux)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, loggedMux))
}
