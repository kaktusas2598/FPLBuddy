package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// ManagerFixtureAnalysis handles requests from frontend for the selected manager's squad fixture analysis
// for selected gameweeks
func ManagerFixtureAnalysis(w http.ResponseWriter, r *http.Request)  {
	managerID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid manager ID", http.StatusBadRequest)
		return
	}

	squadGameweek, err := strconv.Atoi(r.URL.Query().Get("gameweek"))
	if err != nil {
		http.Error(w, "Missing gameweek query param", http.StatusBadRequest)
		return
	}

	startGameweek, err := strconv.Atoi(r.URL.Query().Get("start"))
	if err != nil {
		http.Error(w, "Missing start query param", http.StatusBadRequest)
		return
	}


	gameweekCount, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		http.Error(w, "Missing count query param", http.StatusBadRequest)
		return
	}


	analysis, err := fpl.GetManagerFixtureAnalysis(r.Context(), managerID, squadGameweek, startGameweek, gameweekCount)
	if err != nil {
		log.Printf("failed to calculate the fixture analysis for the manager: %v", err)
		http.Error(w, "Failed fetching the squad picks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(analysis); err != nil {
		log.Printf("Failed to encode the squad fixture analysis: %v", err)
		http.Error(w, "Failed to encode the squad fixture analysis", http.StatusInternalServerError)
	}
}
