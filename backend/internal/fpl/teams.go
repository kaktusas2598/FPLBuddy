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


