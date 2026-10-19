package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetPlayer fetches the complete FPL player list and returns the player which matches the provided id.
// Returns a zero-value Player in case of error or if lookup fails
func GetPlayer(ctx context.Context, id int) (Player, error) {

	data, err := doRequest(ctx, http.MethodGet, baseURL + "elements/", nil)
	if err != nil {
		return Player{}, fmt.Errorf("fetch player %d: %w", id, err)
	}

	var players []Player

	// Parses json encoded bytes slice and stores result in the players slice
	err = json.Unmarshal(data, &players)
	if err != nil {
		return Player{}, fmt.Errorf("failed to decode api/elements response: %w", err)
	}

	for _, player := range players {
		if player.ID == id {
			return player, nil
		}
	}

	return Player{}, fmt.Errorf("failed to find the player id %d", id)
}

