package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jsp220/prayer-tracker/backend/database"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// Item struct represents a document in MongoDB
type Prayer struct {
	Id    		string `json:"id,omitempty" bson:"id,omitempty"`
	CreatedAt 	int64 `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	CreatorId 	string `json:"creatorId" bson:"creatorId"`
	Content 	string `json:"content" bson:"content"`
}

// InsertPrayerHandler - Insert data into MongoDB
func InsertPrayerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var prayer Prayer
	if err := json.NewDecoder(r.Body).Decode(&prayer); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if prayer.Id != "" {
		http.Error(w, "ID should not be provided", http.StatusBadRequest)
		return
	}
	if prayer.CreatedAt != 0 {
		http.Error(w, "CreatedAt should not be provided", http.StatusBadRequest)
		return
	}
	
	prayer.Id = uuid.NewString()
	prayer.CreatedAt = time.Now().UnixMilli()

	collection := database.GetCollection("prayer-tracker-db", "prayers")

	_, err := collection.InsertOne(context.TODO(), prayer)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to insert item: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(prayer)
}

// GetPrayersHandler - Fetch data from MongoDB
func GetPrayersHandler(w http.ResponseWriter, r *http.Request) {
	creatorId := r.URL.Query().Get("creatorId")
	if creatorId == "" {
		http.Error(w, "creatorId is required", http.StatusBadRequest)
		return
	}
	
	collection := database.GetCollection("prayer-tracker-db", "prayers")

	cursor, err := collection.Find(context.TODO(), bson.M{"creatorId": creatorId})
	if err != nil {
		http.Error(w, "Failed to fetch prayers", http.StatusInternalServerError)
		return
	}

	var prayers []Prayer
	if err := cursor.All(context.TODO(), &prayers); err != nil {
		http.Error(w, "Failed to decode prayers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prayers)
}
