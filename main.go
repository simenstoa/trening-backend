package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// our main function
func main() {
	log.Println("Starting trening-backend API on port 8080")
	router := mux.NewRouter()
	router.Use(commonMiddleware)
	router.HandleFunc("/activities", getActivities).Methods("GET")
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	log.Fatal(http.ListenAndServe(":8080", loggedRouter))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getActivities(w http.ResponseWriter, r *http.Request) {
	session, err := NewSession()
	if err != nil {
		log.Fatalf("Unable to connect to mongo: %s", err)
	}
	var activities []activity
	session.GetCollection("trening-s33", "activities").Find(nil).All(&activities)
	json.NewEncoder(w).Encode(activities)
}

type activity struct {
	ID                 string  `json:"id,omitempty" bson:"id,omitempty"`
	Name               string  `json:"name,omitempty" bson:"name,omitempty"`
	Distance           float32 `json:"distance,omitempty" bson:"distance,omitempty"`
	MovingTime         int     `json:"moving_time,omitempty" bson:"moving_time,omitempty"`
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty" bson:"total_elevation_gain,omitempty"`
	HasHeartrate       bool    `json:"has_heartrate,omitempty" bson:"has_heartrate,omitempty"`
	WorkoutType        int     `json:"workout_type,omitempty" bson:"workout_type,omitempty"`
	Type               string  `json:"type,omitempty" bson:"type,omitempty"`
	StartDateLocal     string  `json:"start_date_local,omitempty" bson:"start_date_local,omitempty"`
}
