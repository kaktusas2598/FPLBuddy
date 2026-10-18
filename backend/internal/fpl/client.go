// Package fpl provides a client for interacting with the official Fantasy Premier League API.
package fpl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"io"
)

// Official FPL API base URL
const baseURL = "https://fantasy.premierleague.com/api/"
// Unofficial experimental GraphQL based FPL API base URL
const graphQlURL = "https://fpl-api-6h0d.onrender.com/graphql"

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

// GetTeams Uses GraphQL unofficial API to fetch teams, experimental
func GetTeams(ctx context.Context) ([]Team, error) {
	graphQLQuery := `{"query":"{ teams { items { id name short_name code } } }"}`

	// Creating POST request with graphql query as a body
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, graphQlURL, bytes.NewBufferString(graphQLQuery))
	if err != nil {
		return []Team{}, fmt.Errorf("error creating GraphQL request: %w", err)
	}

	// Actually sending off the request
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return []Team{}, fmt.Errorf("error making GraphQL request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []Team{}, fmt.Errorf("response status is not 200: %d", response.StatusCode)
	}

	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return []Team{}, fmt.Errorf("failed io reading the response: %w", err)
	}

	var teams []Team
	// Creating a nested anonymous struct to allow unmarshalling raw json response into typed structure
	var teamsResult struct {
		Data struct {
			Teams struct {
				Items []Team `json:"items"`
			} `json:"teams"`
		} `json:"data"`
	}
	err = json.Unmarshal(rawResponse, &teamsResult)
	if err != nil {
		return []Team{}, fmt.Errorf("json unmarshal error: %w", err)
	}

	teams = teamsResult.Data.Teams.Items

	return teams, nil
}

// GetCurrentGameweek Uses GraphQL unofficial API to fetch the information about current gameweek, experimental
func GetCurrentGameweek(ctx context.Context) (GameWeek, error) {
	graphQLQuery := `{"query":"{ events(where: { is_current: { eq: true } }) { items { id name deadline_time most_selected { web_name } top_element { web_name } } } }"}`

	// Creating POST request with graphql query as a body
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, graphQlURL, bytes.NewBufferString(graphQLQuery))
	if err != nil {
		return GameWeek{}, fmt.Errorf("error creating GraphQL request: %w", err)
	}

	// Actually sending off the request
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return GameWeek{}, fmt.Errorf("error making GraphQL request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return GameWeek{}, fmt.Errorf("response status is not 200: %d", response.StatusCode)
	}

	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return GameWeek{}, fmt.Errorf("failed io reading the response: %w", err)
	}

	var currentGameweek GameWeek
	// Creating a nested anonymous struct to allow unmarshalling raw json response into typed structure
	var eventResult struct {
		Data struct {
			Events struct {
				Items []GameWeek `json:"items"`
			} `json:"events"`
		} `json:"data"`
	}

	err = json.Unmarshal(rawResponse, &eventResult)
	if err != nil {
		return GameWeek{}, fmt.Errorf("json unmarshal error: %w", err)
	}

	currentGameweek = eventResult.Data.Events.Items[0]

	return currentGameweek, nil
}

