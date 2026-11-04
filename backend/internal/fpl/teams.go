package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTeams fetches the information about all FPL teams
// Returns a nil in case of error or if lookup fails
func GetTeams(ctx context.Context) ([]Team, error) {

	data, err := doRequest(ctx, http.MethodGet, baseURL + "bootstrap-static", nil)
	if err != nil {
		return nil, fmt.Errorf("fetch teams: %w", err)
	}

	// Custom type for extracting teams
	var bootstrapResult struct {
		Teams []Team `json:"teams"`
	}


	// Parses json encoded bytes slice and stores result in the custom result type
	err = json.Unmarshal(data, &bootstrapResult)
	if err != nil {
		return nil, fmt.Errorf("failed to decode api/bootstrap-static response: %w", err)
	}

	return bootstrapResult.Teams, nil
}

