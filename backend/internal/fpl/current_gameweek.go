package fpl

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

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

