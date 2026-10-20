package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSquad fetches the squad selection with player information for selected manager and gameweek
// Returns a zero-value SquadDetails in case of error or if lookup fails
func GetSquad(ctx context.Context, managerID int, gameWeekID int) (SquadDetails, error) {

	url := fmt.Sprintf("%sentry/%d/event/%d/picks", baseURL, managerID, gameWeekID)
	data, err := doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return SquadDetails{}, fmt.Errorf("fetch squad: %w", err)
	}

	var squad Squad
	// Parses json encoded bytes slice and stores result in the custom result type
	err = json.Unmarshal(data, &squad)
	if err != nil {
		return SquadDetails{}, fmt.Errorf("failed to decode entry event picks response: %w", err)
	}

	// NOTE: this is not ideal probably at the moment, so might want to change later, cache or sth
	data, err = doRequest(ctx, http.MethodGet, baseURL + "elements/", nil)
	if err != nil {
		return SquadDetails{}, fmt.Errorf("fetch players: %w", err)
	}

	var squadDetails SquadDetails
	var players []Player

	// Parses json encoded bytes slice and stores result in the players slice
	err = json.Unmarshal(data, &players)
	if err != nil {
		return SquadDetails{}, fmt.Errorf("failed to decode api/elements response: %w", err)
	}

	playersByID := make(map[int]Player, len(players))
	for _, player := range players {
		playersByID[player.ID] = player
	}

	for _, pick := range squad.Picks {
		player, ok := playersByID[pick.PlayerID]
		if !ok {
			return SquadDetails{}, fmt.Errorf("player %d not found", pick.PlayerID)
		}

		squadDetails.Picks = append(squadDetails.Picks, SquadPlayer{
			Player: player,
			Pick: pick,
		})
	}

	return squadDetails, nil
}

