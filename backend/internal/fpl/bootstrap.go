package fpl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// BootstrapData is a custom data type defining the response from FPL API's /api/bootstrap-static endpoint
type BootstrapData struct {
	Events  []GameWeek `json:"events"`
	Players []Player   `json:"elements"`
	Teams   []Team     `json:"teams"`
}

// bootstrapCache is internal structure used for caching bootstrap data in the memory
var bootstrapCache struct {
	sync.Mutex
	data      BootstrapData
	expiresAt time.Time
}

// GetBootstrapData fetches all teams, players and event information from official FPL API's bootstrap-static
// endpoint to be reused by other FPLBuddy Go endpoints which is then cached to avoid retrieving this data multi times
func GetBootstrapData(ctx context.Context) (BootstrapData, error) {

	// Prevent another routine from accessing data while this is being called and then ensure lock is released
	bootstrapCache.Lock()
	defer bootstrapCache.Unlock()

	// Return cached data if not expired
	if time.Now().Before(bootstrapCache.expiresAt) {
		return bootstrapCache.data, nil
	}

	data, err := doRequest(ctx, http.MethodGet, baseURL + "bootstrap-static", nil)
	if err != nil {
		return BootstrapData{}, fmt.Errorf("fetch bootstrap data: %w", err)
	}

	var bootstrapData BootstrapData
	// Parses json encoded bytes slice and stores result in the custom result type
	err = json.Unmarshal(data, &bootstrapData)
	if err != nil {
		return BootstrapData{}, fmt.Errorf("failed to decode api/bootstrap-static response: %w", err)
	}

	// Set cache and cache expiry of 5 minutes
	bootstrapCache.data = bootstrapData
	bootstrapCache.expiresAt = time.Now().Add(5 * time.Minute)

	return bootstrapData, nil
}
