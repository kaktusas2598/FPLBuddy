package fpl

import "testing"

func TestAnalyseFixtureDifficulty(t *testing.T) {
	gameweek := 9

	fixtures := []Fixture{
		{
		GameWeek: &gameweek,
		TeamH: 1,
		TeamA: 2,
		TeamHDifficulty: 2,
		TeamADifficulty: 4,
		},
	}

	difficulty, matchingFixtures := AnalyseFixtureDifficulty(1, fixtures, 9, 1)

	if difficulty != 2 {
    t.Errorf("expected difficulty 2, got %v", difficulty)
}

	if len(matchingFixtures) != 1 {
		t.Errorf("expected 1 matching fixture, got %d", len(matchingFixtures))
	}
}

// TODO: Add more Tests here with different assertions
