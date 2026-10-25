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

	result := AnalyzeFixtureDifficulty(1, fixtures, 9, 1)

	if result != 2 {
		t.Errorf("expected difficulty 2, got %v", result)
	}
}

// TODO: Add more Tests here with different assertions
