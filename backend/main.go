package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Hello World",
		Data:    "1",
	}
	json.NewEncoder(w).Encode(response)
}

func AmyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Hi Amy!",
		Data:    "2",
	}
	json.NewEncoder(w).Encode(response)
}

func EmmaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Hi Emma!",
		Data:    "3",
	}
	json.NewEncoder(w).Encode(response)
}

func JayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Hi Jay!",
		Data:    "4",
	}
	json.NewEncoder(w).Encode(response)
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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Wrap router with CORS middleware
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/amy", AmyHandler)
	mux.HandleFunc("/emma", EmmaHandler)
	mux.HandleFunc("/jay", JayHandler)
	handler := enableCORS(mux)

	// Start the server
	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
