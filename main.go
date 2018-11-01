package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc("/activities", getActivities).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getActivities(w http.ResponseWriter, r *http.Request) {
	var activities []activity
	activities = append(activities, activity{ID: "1", Name: "3x8 på mølla", Distance: 5856.9, MovingTime: 2162, Type: "Run", WorkoutType: 0, StartDateLocal: "2018-10-17T16:49:38Z"})
	activities = append(activities, activity{ID: "2", Name: "3x8 på mølla", Distance: 5856.9, MovingTime: 2162, Type: "Run", WorkoutType: 0, StartDateLocal: "2018-10-17T16:49:38Z"})
	json.NewEncoder(w).Encode(activities)
}

type activity struct {
	ID                 string  `json:"id,omitempty"`
	Name               string  `json:"name,omitempty"`
	Distance           float32 `json:"distance,omitempty"`
	MovingTime         int     `json:"moving_time,omitempty"`
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`
	HasHeartrate       bool    `json:"has_heartrate,omitempty"`
	WorkoutType        int     `json:"workout_type,omitempty"`
	Type               string  `json:"type,omitempty"`
	StartDateLocal     string  `json:"start_date_local,omitempty"`
}
