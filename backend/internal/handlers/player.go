package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

func Player(w http.ResponseWriter, r *http.Request)  {
    // Parse and validate the player ID
	playerID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid player ID", http.StatusBadRequest)
		return
	}

	player, err := fpl.GetPlayer(r.Context(), playerID)
	if err != nil {
		// TODO: improve error handling here, could be all kinds of errors
		http.Error(w, "Player Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(player); err != nil {
		log.Printf("Failed to encode player: %v", err)
		//http.Error(w, "Player Not Found", http.StatusInternalServerError)
	}
}
