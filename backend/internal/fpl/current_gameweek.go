package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCurrentGameweek(ctx context.Context) (GameWeek, error) {

	data, err := doRequest(ctx, http.MethodGet, baseURL + "bootstrap-static", nil)
	if err != nil {
		return GameWeek{}, fmt.Errorf("fetch current gameweek: %w", err)
	}

	var bootstrapStaticResult struct {
		Events []GameWeek `json:"events"`
	}

	// Parses json encoded bytes slice and stores result in the custom result type
	err = json.Unmarshal(data, &bootstrapStaticResult)
	if err != nil {
		return GameWeek{}, fmt.Errorf("failed to decode api/bootstrap-static response: %w", err)
	}

	for _, gameWeek := range bootstrapStaticResult.Events {
		if gameWeek.IsCurrent {
			return gameWeek, nil
		}
	}

	return GameWeek{}, fmt.Errorf("failed to find the current gameweek")
}

