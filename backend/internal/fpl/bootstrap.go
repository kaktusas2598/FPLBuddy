package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// BootstrapData is a custom data type defining the response from FPL API's /api/bootstrap-static endpoint
type BootstrapData struct {
	Events  []GameWeek `json:"events"`
	Players []Player   `json:"elements"`
	Teams   []Team     `json:"teams"`
}

// GetBootstrapData fetches all teams, players and event information from official FPL API's bootstrap-static
// endpoint to be reused by other FPLBuddy Go endpoints which can be later cached
func GetBootstrapData(ctx context.Context) (BootstrapData, error) {

	data, err := doRequest(ctx, http.MethodGet, baseURL + "bootstrap-static", nil)
	if err != nil {
		return BootstrapData{}, fmt.Errorf("fetch bootstrap data: %w", err)
	}

	var bootstrapData BootstrapData
	// Parses json encoded bytes slice and stores result in the custom result type
	err = json.Unmarshal(data, &bootstrapData)
	if err != nil {
		return BootstrapData{}, fmt.Errorf("failed to decode api/bootstrap-static response: %w", err)
	}

	return bootstrapData, nil
}
