package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetManager fetches the information about an FPL manager by ID.
func GetManager(ctx context.Context, id int) (json.RawMessage, error) {

	url := fmt.Sprintf("%sentry/%d/", baseURL, id)
	data, err := doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("fetch manager %d: %w", id, err)
	}

	return json.RawMessage(data), nil
}

