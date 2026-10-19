// Package handlers implements the HTTP endpoints exposed by FPLBuddy.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// Player handles requests for the requested prem player information
func Player(w http.ResponseWriter, r *http.Request)  {
    // Parse and validate the player ID
	playerID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Failed fetching the player", http.StatusInternalServerError)
		return
	}

	player, err := fpl.GetPlayer(r.Context(), playerID)
	if err != nil {
		log.Printf("failed to fetch player %d: %v", playerID, err)
		http.Error(w, "Player Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(player); err != nil {
		http.Error(w, "Failed to encode the player", http.StatusInternalServerError)
	}
}
