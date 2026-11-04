package fpl

import (
	"context"
	"fmt"
)

// GetPlayers fetches the information about all FPL players
func GetPlayers(ctx context.Context) ([]Player, error) {

	bootstrapData, err := GetBootstrapData(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch players: %w", err)
	}

	return bootstrapData.Players, nil

}

