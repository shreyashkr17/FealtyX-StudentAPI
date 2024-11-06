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

	// Serve index.html for the root path
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Add existing student handler
	mux.HandleFunc("/students/", handlers.StudentHandler)

	// Wrap with logging middleware
	loggedMux := middleware.LoggingMiddleware(mux)

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, loggedMux))
}
