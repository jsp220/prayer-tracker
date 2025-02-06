package routes

import (
	"net/http"

	"github.com/jsp220/prayer-tracker/backend/controllers"
)

// RegisterRoutes sets up the API routes
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.HomeHandler) 
	mux.HandleFunc("/users", controllers.GetUsersHandler)   // GET all users
	mux.HandleFunc("/users/insert", controllers.InsertUserHandler) // POST insert user
	mux.HandleFunc("/prayers", controllers.GetPrayersHandler)   // GET all prayers
	mux.HandleFunc("/prayers/insert", controllers.InsertPrayerHandler) // POST insert prayer
}
