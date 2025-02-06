package controllers

import (
	"encoding/json"
	"net/http"
)

// HomeHandler - Default API Response
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to Prayer Tracker API!"})
}