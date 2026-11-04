package fpl

import (
	"context"
	"fmt"
)

// GetCurrentGameweek fetches the information about current gameweek
// Returns a zero-value GameWeekDetails in case of error or if lookup fails
func GetCurrentGameweek(ctx context.Context) (GameWeekDetails, error) {

	bootstrapData, err := GetBootstrapData(ctx)
	if err != nil {
		return GameWeekDetails{}, fmt.Errorf("fetch current gameweek: %w", err)
	}

	var currentGameweekDetails GameWeekDetails
	var currentGameweek GameWeek
	for _, gameWeek := range bootstrapData.Events {
		if gameWeek.IsCurrent {
			currentGameweek = gameWeek
		}
	}

	if currentGameweek.ID == 0 {
		return GameWeekDetails{}, fmt.Errorf("failed to find the current gameweek")
	}

	var mostPointsPlayer Player
	var topSelectedPlayer Player
	for _, player := range bootstrapData.Players {
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

