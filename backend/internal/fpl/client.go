// Package fpl provides a client for interacting with the official Fantasy Premier League API.
package fpl

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Official FPL API base URL
const baseURL = "https://fantasy.premierleague.com/api/"
// Unofficial experimental GraphQL based FPL API base URL
const graphQlURL = "https://fpl-api-6h0d.onrender.com/graphql"

// TODO: in the near future, I want to introduce type for this!
var httpClient = &http.Client{
	Timeout: 30 * time.Second, // For now timeout is long because unofficial GraphQL can take some time due to cold starts
}

// doRequest is internal helper to create and execute HTTP request
func doRequest(ctx context.Context, method, url string, body io.Reader) ([]byte, error) {

	request, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating the %s request for %s: %w", method, url, err)
	}

	// NOTE: For now assuming any POST request will be made to GraphQL unoffical FPL API
	if method == http.MethodPost {
		request.Header.Set("Content-Type", "application/json")
	}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error making the %s request for %s: %w", method, url, err)
	}

	defer response.Body.Close()

	// NOTE: graphql responses will still be 200 in case of errors and error will be returned in json response!
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response status for %s is not 200: %d", url, response.StatusCode)
	}

	rawResponse, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s response: %w", url, err)
	}

	return rawResponse, nil
}

