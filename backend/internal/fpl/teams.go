package fpl

import (
	"context"
	"fmt"
)

// GetTeams fetches the information about all FPL teams
func GetTeams(ctx context.Context) ([]Team, error) {

	bootstrapData, err := GetBootstrapData(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch teams: %w", err)
	}

	return bootstrapData.Teams, nil
}

