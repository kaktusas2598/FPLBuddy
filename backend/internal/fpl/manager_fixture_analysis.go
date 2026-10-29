package fpl

import (
	"context"
	"fmt"
)

// GetManagerFixtureAnalysis calculates the average fixture difficulty of each squad pick for selected manager and
// selected gameweek. Returns nil in case of error or if calculation fails.
func GetManagerFixtureAnalysis(ctx context.Context, managerID int, squadGameweek int, startGameweek int, gameweekCount int) ([]PlayerFixtureAnalysis, error) {

	squadDetails, err := GetSquad(ctx, managerID, squadGameweek)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch manager squad while trying to calculate the fixture analysis: %w", err)
	}

	fixtures, err := GetFixtures(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch fixtures while trying to calculate the manager's fixture analysis: %w", err)
	}

	return AnalyseSquadFixtures(squadDetails, fixtures, startGameweek, gameweekCount), nil
}
