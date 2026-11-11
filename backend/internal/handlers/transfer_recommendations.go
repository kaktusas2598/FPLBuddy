package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// TransferRecommendations handles requests from frontend to analyse managers squad for selected gameweek
// and propose recommended transfers
func TransferRecommendations(w http.ResponseWriter, r *http.Request)  {
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

	if squadGameweek < 1 || startGameweek < 1 || gameweekCount < 1 {
		http.Error(w, "Gameweek parameters must be positive", http.StatusBadRequest)
		return
	}

	squad, err := fpl.GetSquad(r.Context(), managerID, squadGameweek)
	if err != nil {
		http.Error(w, "Failed to fetch squad for transfer recommendations", http.StatusInternalServerError)
		return
	}

	players, err := fpl.GetPlayers(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch players for transfer recommendations", http.StatusInternalServerError)
		return
	}

	fixtures, err := fpl.GetFixtures(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch fixtures for transfer recommendations", http.StatusInternalServerError)
		return
	}

	manager, err := fpl.GetManager(r.Context(), managerID)
	if err != nil {
		http.Error(w, "Failed to fetch manager information for transfer recommendations", http.StatusInternalServerError)
		return
	}

	var managerInfo struct {
		LastDeadlineBank int `json:"last_deadline_bank"`
	}

	if err := json.Unmarshal(manager, &managerInfo); err != nil {
		http.Error(w, "Failed to decode manager information for transfer recommendations", http.StatusInternalServerError)
		return
	}

	transferRecommendations := fpl.RecommendTransfers(
		squad,
		players,
		fixtures,
		startGameweek,
		gameweekCount,
		managerInfo.LastDeadlineBank,
		)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(transferRecommendations); err != nil {
		log.Printf("Failed to encode the manager transfer recommendations: %v", err)
		http.Error(w, "Failed to encode the manager transfer recommendations", http.StatusInternalServerError)
	}
}
