package fpl

// AnalyzeFixtureDifficulty calculates the fixture difficulty for selected team
// and selected number for upcoming gameweeeks
func AnalyzeFixtureDifficulty(
    teamID int,
    fixtures []Fixture,
    nextGameweek int,
    gameweeks int,
) float64  {

	var fixturesCount, sum int
	for _, fixture := range fixtures {
		if (fixture.GameWeek == nil) {
			continue
		}

		// TODO: in the future validate inputs (like negative number of weeks passed in) and account for finished weeks
		// Filter fixtures for correct gameweeks
		if *fixture.GameWeek >= nextGameweek && *fixture.GameWeek < nextGameweek + gameweeks {
			// Filter fixtures for correct teams
			if fixture.TeamA == teamID || fixture.TeamH == teamID {
				fixturesCount++
				if fixture.TeamA == teamID {
					sum += fixture.TeamADifficulty
				} else {
					sum += fixture.TeamHDifficulty
				}
			}
		}
	}

	if fixturesCount == 0 {
		return 0
	}

	// Most basic averaging for now
	return (float64(sum) / float64(fixturesCount))

}
