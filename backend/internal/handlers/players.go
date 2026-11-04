package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// Players handles requests from frontend for the premier league players information
func Players(w http.ResponseWriter, r *http.Request)  {
	players, err := fpl.GetPlayers(r.Context())
	if err != nil {
		log.Printf("failed to fetch the players: %v", err)
		http.Error(w, "Failed fetching the players", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(players); err != nil {
		log.Printf("Failed to encode the players: %v", err)
		http.Error(w, "Failed to encode the players", http.StatusInternalServerError)
	}
}
