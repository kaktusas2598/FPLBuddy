package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// PlayerSummary handles requests from the frontend for the requested prem player summary information
func PlayerSummary(w http.ResponseWriter, r *http.Request)  {
    // Parse and validate the player ID
	playerID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Failed fetching the player summary", http.StatusInternalServerError)
		return
	}

	player, err := fpl.GetPlayerSummary(r.Context(), playerID)
	if err != nil {
		log.Printf("failed to fetch player summary %d: %v", playerID, err)
		http.Error(w, "Player Summary Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(player); err != nil {
		http.Error(w, "Failed to encode the player summary", http.StatusInternalServerError)
	}
}
