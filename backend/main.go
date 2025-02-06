package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jsp220/prayer-tracker/backend/database"
	"github.com/jsp220/prayer-tracker/backend/routes"

	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

// CORS Middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Allow frontend
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	database.ConnectMongoDB()
	defer database.DisconnectMongoDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	// Register routes
	routes.RegisterRoutes(mux)

	// Wrap router with CORS middleware
	handler := enableCORS(mux)

	// Start the server
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
