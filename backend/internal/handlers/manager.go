package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// Manager handles requests for the requested FPL manager information
func Manager(w http.ResponseWriter, r *http.Request)  {
    // Parse and validate the manager/squad ID
	managerID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Failed fetching the manager", http.StatusInternalServerError)
		return
	}

	manager, err := fpl.GetManager(r.Context(), managerID)
	if err != nil {
		http.Error(w, "Manager Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(manager); err != nil {
		http.Error(w, "Failed to encode the manager", http.StatusInternalServerError)
	}
}
