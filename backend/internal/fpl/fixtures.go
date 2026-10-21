package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetFixtures fetches the complete FPL fixture list
// Returns a zero-value Fixture slice in case of error or if lookup fails
func GetFixtures(ctx context.Context) ([]Fixture, error) {

	data, err := doRequest(ctx, http.MethodGet, baseURL + "fixtures/", nil)
	if err != nil {
		return []Fixture{}, fmt.Errorf("fetch fixtures: %w", err)
	}

	var fixtures []Fixture

	// Parses json encoded bytes slice and stores result in the fixtures slice
	err = json.Unmarshal(data, &fixtures)
	if err != nil {
		return []Fixture{}, fmt.Errorf("failed to decode api/fixtures response: %w", err)
	}

	return fixtures, nil
}

