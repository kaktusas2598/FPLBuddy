package fpl

// AnalyseFixtureDifficulty calculates the fixture difficulty for selected team
// and selected number for upcoming gameweeeks
func AnalyseFixtureDifficulty(
	teamID int,
	fixtures []Fixture,
	nextGameweek int,
	gameweeks int,
) float64 {

	var fixturesCount, sum int
	for _, fixture := range fixtures {
		if fixture.GameWeek == nil {
			continue
		}

		// TODO: in the future validate inputs (like negative number of weeks passed in) and account for finished weeks
		// Filter fixtures for correct gameweeks
		if *fixture.GameWeek >= nextGameweek && *fixture.GameWeek < nextGameweek+gameweeks {
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

// AnalyseSquadFixtures calculates average fixture difficulty for the whole selected squad based on specified
// number of weeks starting from nextGameweek
// TODO: Write Unit test for this!
func AnalyseSquadFixtures(
	squad SquadDetails,
	fixtures []Fixture,
	nextGameweek int,
	gameweeks int,
) []PlayerFixtureAnalysis {

	var fixtureAnalysis []PlayerFixtureAnalysis
	for _, player := range squad.Picks {
		fixtureDifficulty := AnalyseFixtureDifficulty(player.Player.Team, fixtures, nextGameweek, gameweeks)
		fixtureAnalysis = append(fixtureAnalysis, PlayerFixtureAnalysis{
			PlayerID: player.Player.ID,
			AverageDifficulty: fixtureDifficulty})
	}

	return fixtureAnalysis
}
