package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jsp220/prayer-tracker/backend/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Item struct represents a document in MongoDB
type User struct {
	Id    		string `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt 	int64  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	Name  		string `json:"name" bson:"name"`
	Email 		string `json:"email" bson:"email"`
}

// InsertUserHandler - Insert data into MongoDB
func InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if user.Id != "" {
		http.Error(w, "ID should not be provided", http.StatusBadRequest)
		return
	}
	
	user.Id = uuid.NewString()

	collection := database.GetCollection("prayer-tracker-db", "users")

	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert item: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetItemsHandler - Fetch data from MongoDB
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	collection := database.GetCollection("prayer-tracker-db", "users")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	var users []User
	if err := cursor.All(context.TODO(), &users); err != nil {
		http.Error(w, "Failed to decode users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
