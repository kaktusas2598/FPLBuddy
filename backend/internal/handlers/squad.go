package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// Squad handles requests from frontend for the selected manager and gameweek's squad picks
func Squad(w http.ResponseWriter, r *http.Request)  {
	managerID, err := strconv.Atoi(r.PathValue("manager_id"))
	if err != nil {
		http.Error(w, "Invalid manager ID", http.StatusBadRequest)
		return
	}

	gameweekID, err := strconv.Atoi(r.PathValue("gameweek_id"))
	if err != nil {
		http.Error(w, "Invalid gameweek ID", http.StatusBadRequest)
		return
	}

	squad, err := fpl.GetSquad(r.Context(), managerID, gameweekID)
	if err != nil {
		log.Printf("failed to fetch the squad picks: %v", err)
		http.Error(w, "Failed fetching the squad picks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(squad); err != nil {
		log.Printf("Failed to encode the squad picks: %v", err)
		http.Error(w, "Failed to encode the squad picks", http.StatusInternalServerError)
	}
}
