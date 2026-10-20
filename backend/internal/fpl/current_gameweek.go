package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCurrentGameweek(ctx context.Context) (GameWeekDetails, error) {

	data, err := doRequest(ctx, http.MethodGet, baseURL + "bootstrap-static", nil)
	if err != nil {
		return GameWeekDetails{}, fmt.Errorf("fetch current gameweek: %w", err)
	}

	// Custom type for extracting both events/gameweeks and elements/players
	var bootstrapResult struct {
		Events []GameWeek `json:"events"`
		Elements []Player `json:"elements"`
	}

	// Parses json encoded bytes slice and stores result in the custom result type
	err = json.Unmarshal(data, &bootstrapResult)
	if err != nil {
		return GameWeekDetails{}, fmt.Errorf("failed to decode api/bootstrap-static response: %w", err)
	}

	var currentGameweekDetails GameWeekDetails
	var currentGameweek GameWeek
	for _, gameWeek := range bootstrapResult.Events {
		if gameWeek.IsCurrent {
			currentGameweek = gameWeek
		}
	}

	if currentGameweek.ID == 0 {
		return GameWeekDetails{}, fmt.Errorf("failed to find the current gameweek")
	}

	var mostPointsPlayer Player
	var topSelectedPlayer Player
	for _, player := range bootstrapResult.Elements {
		if player.ID == currentGameweek.MostPoints {
			mostPointsPlayer = player
		}
		if player.ID == currentGameweek.MostSelected {
			topSelectedPlayer = player
		}

	}

	currentGameweekDetails.ID = currentGameweek.ID
	currentGameweekDetails.Name = currentGameweek.Name
	currentGameweekDetails.DeadlineTime = currentGameweek.DeadlineTime
	currentGameweekDetails.MostPoints = mostPointsPlayer
	currentGameweekDetails.MostSelected = topSelectedPlayer

	return currentGameweekDetails, nil
}

