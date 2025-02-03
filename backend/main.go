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
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{Message: "Hello World"}
	json.NewEncoder(w).Encode(response)
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

	http.HandleFunc("/", HomeHandler)

	fmt.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
