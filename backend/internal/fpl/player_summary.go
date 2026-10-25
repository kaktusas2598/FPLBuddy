package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetPlayerSummary fetches the past and the upcoming FPL fixtures for provided player ID.
// Returns a zero-value PlayerSummary in case of error or if lookup fails
func GetPlayerSummary(ctx context.Context, id int) (PlayerSummary, error) {

	url := fmt.Sprintf("%selement-summary/%d", baseURL, id)
	data, err := doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return PlayerSummary{}, fmt.Errorf("fetch the summary for player %d: %w", id, err)
	}

	var playerSummary PlayerSummary

	// Parses json encoded bytes slice and stores result in the players slice
	err = json.Unmarshal(data, &playerSummary)
	if err != nil {
		return PlayerSummary{}, fmt.Errorf("failed to decode api/element-summary response: %w", err)
	}

	return playerSummary, nil
}

