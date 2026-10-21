// Package handlers implements the HTTP endpoints exposed by FPLBuddy.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
)

// Fixtures handles requests from the frontend for all the prem fixtures of the current season information
func Fixtures(w http.ResponseWriter, r *http.Request)  {
	fixtures, err := fpl.GetFixtures(r.Context())
	if err != nil {
		log.Printf("failed to fetch the fixtures: %v", err)
		http.Error(w, "Fixtures fetch failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(fixtures); err != nil {
		http.Error(w, "Failed to encode the fixtures", http.StatusInternalServerError)
	}
}
