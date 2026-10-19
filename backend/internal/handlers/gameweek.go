package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// CurrentGameWeek handles requests for the current gameweek information
func CurrentGameWeek(w http.ResponseWriter, r *http.Request)  {
	currentGameWeek, err := fpl.GetCurrentGameweek(r.Context())
	if err != nil {
		log.Printf("failed to fetch the current gameweek: %v", err)
		http.Error(w, "Failed fetching the current gameweek", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(currentGameWeek); err != nil {
		log.Printf("Failed to encode the current gameweek: %v", err)
		http.Error(w, "Failed to encode the current gameweek", http.StatusInternalServerError)
	}
}
