package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetPlayers fetches the information about all FPL players
// Returns a nil in case of error or if lookup fails
func GetPlayers(ctx context.Context) ([]Player, error) {

	data, err := doRequest(ctx, http.MethodGet, baseURL + "bootstrap-static", nil)
	if err != nil {
		return nil, fmt.Errorf("fetch players: %w", err)
	}

	// Custom type for extracting players
	var bootstrapResult struct {
		Players []Player `json:"elements"`
	}


	// Parses json encoded bytes slice and stores result in the custom result type
	err = json.Unmarshal(data, &bootstrapResult)
	if err != nil {
		return nil, fmt.Errorf("failed to decode api/bootstrap-static response: %w", err)
	}

	return bootstrapResult.Players, nil
}

