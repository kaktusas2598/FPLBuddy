package fpl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTeams Uses GraphQL unofficial API to fetch teams, experimental
func GetTeams(ctx context.Context) ([]Team, error) {
	graphQLQuery := `{"query":"{ teams { items { id name short_name code } } }"}`

	data, err := doRequest(ctx, http.MethodPost, graphQlURL, bytes.NewBufferString(graphQLQuery))
	if err != nil {
		return []Team{}, fmt.Errorf("fetch teams: %w", err)
	}

	// Creating a nested anonymous struct to allow unmarshalling raw json response into typed structure
	var teamsResult struct {
		Data struct {
			Teams struct {
				Items []Team `json:"items"`
			} `json:"teams"`
		} `json:"data"`
	}
	err = json.Unmarshal(data, &teamsResult)
	if err != nil {
		return []Team{}, fmt.Errorf("json unmarshal error: %w", err)
	}

	return teamsResult.Data.Teams.Items, nil
}

// GetCurrentGameweek Uses GraphQL unofficial API to fetch the information about current gameweek, experimental
func GetCurrentGameweek(ctx context.Context) (GameWeek, error) {
	graphQLQuery := `{"query":"{ events(where: { is_current: { eq: true } }) { items { id name deadline_time most_selected { web_name } top_element { web_name } } } }"}`

	data, err := doRequest(ctx, http.MethodPost, graphQlURL, bytes.NewBufferString(graphQLQuery))
	if err != nil {
		return GameWeek{}, fmt.Errorf("fetch current gameweek: %w", err)
	}

	// Creating a nested anonymous struct to allow unmarshalling raw json response into typed structure
	var eventResult struct {
		Data struct {
			Events struct {
				Items []GameWeek `json:"items"`
			} `json:"events"`
		} `json:"data"`
	}

	err = json.Unmarshal(data, &eventResult)
	if err != nil {
		return GameWeek{}, fmt.Errorf("json unmarshal error: %w", err)
	}

	if len(eventResult.Data.Events.Items) == 0 {
		return GameWeek{}, fmt.Errorf("no current gameweek found")
	}

	return eventResult.Data.Events.Items[0], nil
}

