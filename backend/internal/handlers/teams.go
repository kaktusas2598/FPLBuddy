package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// Teams handles requests from frontend for the premier league teams information
func Teams(w http.ResponseWriter, r *http.Request)  {
	teams, err := fpl.GetTeams(r.Context())
	if err != nil {
		log.Printf("failed to fetch the teams: %v", err)
		http.Error(w, "Failed fetching the teams", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(teams); err != nil {
		log.Printf("Failed to encode the teams: %v", err)
		http.Error(w, "Failed to encode the teams", http.StatusInternalServerError)
	}
}
