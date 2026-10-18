package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

func CurrentGameWeek(w http.ResponseWriter, r *http.Request)  {
	currentGameWeek, err := fpl.GetCurrentGameweek(r.Context())
	if err != nil {
		// TODO: improve error handling here, could be all kinds of errors
		//http.Error(w, "Errors f", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(currentGameWeek); err != nil {
		log.Printf("Failed to encode the current gameweek: %v", err)
		//http.Error(w, "Player Not Found", http.StatusInternalServerError)
	}
}
