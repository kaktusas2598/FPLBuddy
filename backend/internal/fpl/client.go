// Package fpl provides a client for interacting with the official Fantasy Premier League API.
package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Oficcial FPL API base URL
const baseURL = "https://fantasy.premierleague.com/api/"

// GetPlayer fetches the complete FPL player list and returns the player which matches the provided id.
// Returns a zero-value Player in case of error or if lookup fails
func GetPlayer(ctx context.Context, id int) (Player, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL + "elements/", nil)
	if err != nil {
		return Player{}, fmt.Errorf("error creating FPL API request: %w", err)
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return Player{}, fmt.Errorf("error making FPL API request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return Player{}, fmt.Errorf("response status is not 200: %d", response.StatusCode)
	}

	var players []Player

	// Decode reads JSON directly from HTTP response body into a slice of Player structs.
	// The decoder matches JSON fields in a struct useing their `json` tags.
	err = json.NewDecoder(response.Body).Decode(&players)
	if err != nil {
		return Player{}, fmt.Errorf("failed to decode api/elements response: %w", err)
	}

	// "_" is used to discard the numerical index, could also be used as for index, player
	for _, player := range players {
		if player.ID == id {
			return player, nil
		}
	}

	return Player{}, fmt.Errorf("failed to find the player id %d", id)
}
