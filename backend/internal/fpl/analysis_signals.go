package fpl

import "strconv"

// EvaluatePlayerAvailability determines if provided player is available for a match and returns a new PlayerSignal
func EvaluatePlayerAvailability(player Player) *PlayerSignal {

	var playerSignal *PlayerSignal
	reason := player.News

	// NOTE: In go break statements are not needed to avoid fall through to the next case!
	switch player.Status {
	case "a": // Available
		return nil
	case "d": // Doubtful
		if reason == "" {
			reason = "Player is currently doubtful"
		}

		playerSignal = &PlayerSignal{
			PlayerID: player.ID,
			Type:     SignalAvailabilityRisk,
			Severity: SeverityMedium,
			Reason:   reason,
		}
	case "i", "s", "u": // Injured, Suspended or Unavailable
		if reason == "" {
			reason = "Player is currently injured, suspended or unavailable"
		}

		playerSignal = &PlayerSignal{
			PlayerID: player.ID,
			Type:     SignalAvailabilityRisk,
			Severity: SeverityHigh,
			Reason:   reason,
		}
	default:
	}

	return playerSignal
}

// EvaluatePlayerForm determines if provided player had a good form in the past 30 days
func EvaluatePlayerForm(player Player) *PlayerSignal {

	var playerSignal *PlayerSignal

	playerForm, err := strconv.ParseFloat(player.Form, 64)
	if err != nil {
		// TODO: handle this better!
		return nil
	}

	// NOTE: this is provisional for now as player.Minutes are minutes for the whole season!
	// No need to account for form if player has barely even played, in the near future we want to account
	// for player.Starts as well obviously
	if player.Minutes < 180 {
		return nil
	}

	if playerForm <= 2.0 {
		playerSignal = &PlayerSignal{
			PlayerID: player.ID,
			Type:     SignalPoorForm,
			Severity: SeverityMedium,
			Reason:   "Player has shown poor form over the last 30 days",
		}
	}

	if playerForm > 2.0 && playerForm <= 3.0 {
		playerSignal = &PlayerSignal{
			PlayerID: player.ID,
			Type:     SignalPoorForm,
			Severity: SeverityLow,
			Reason:   "Player has shown average form over the last 30 days",
		}
	}

	return playerSignal
}

// EvaluatePlayer calculates all types of signals for individual player
func EvaluatePlayer(player Player) []PlayerSignal {

	var playerSignals []PlayerSignal

	playerAvailabilityRiskSignal := EvaluatePlayerAvailability(player)
	if playerAvailabilityRiskSignal != nil {
		playerSignals = append(playerSignals, *playerAvailabilityRiskSignal)
	}

	playerPoorFormSignal := EvaluatePlayerForm(player)
	if playerPoorFormSignal != nil {
		playerSignals = append(playerSignals, *playerPoorFormSignal)
	}

	return playerSignals
}

// EvaluatePlayerFixtures produces a signal of difficult fixtures based on supplied PlayerFixtureAnalysis
func EvaluatePlayerFixtures(analysis PlayerFixtureAnalysis) *PlayerSignal {
	if len(analysis.Fixtures) == 0 {
		return nil
	}

	// NOTE: 3.4 value based on my observations, this is not yet set in stone!!!
	if analysis.AverageDifficulty >= 3.4 {
		return &PlayerSignal{
			PlayerID: analysis.PlayerID,
			Type:     SignalDifficultFixtures,
			Severity: SeverityMedium,
			Reason:   "Player faces difficult upcoming fixtures",
		}
	}

	return nil
}


// EvaluateSquad evaluates all manager's squad picks and produces a slice of player signals
func EvaluateSquad(squad SquadDetails, fixtureAnalysis []PlayerFixtureAnalysis) []PlayerSignal {

	var playerSignals []PlayerSignal

	for _, pick := range squad.Picks {
		playerSignals = append(playerSignals, EvaluatePlayer(pick.Player)...)
	}

	for _, analysis := range fixtureAnalysis {
		if signal := EvaluatePlayerFixtures(analysis); signal != nil {
			playerSignals = append(playerSignals, *signal)
		}
	}

	return playerSignals
}

// IdentifySquadIssues inspects all player signals from selected squads and produces a slice of issues
func IdentifySquadIssues(squad SquadDetails, signals []PlayerSignal) []SquadIssue {

	var issues []SquadIssue

	// To create index of players and count available players by position
	players := make(map[int]Player)
	available := make(map[int]int)

	for _, pick := range squad.Picks {
		player := pick.Player
		players[player.ID] = player
		// Count available players by position: GK, DEF, MID and FWD
		if player.Status == "a" {
			available[player.ElementType]++
		}
	}

	// Minimum players by position required to field a valid starting XI
	minimum := map[int]int{
		1: 1, // Goalkeepers
		2: 3, // Defenders
		3: 2, // Midfielders
		4: 1, // Forwards
	}

	for _, signal := range signals {
		player, exists := players[signal.PlayerID]
		if !exists {
			continue
		}

		var priority SignalSeverity
		var reason string

		// TODO: we should account for all types of signal instead of analysing each one separately!
		// Player with poor form(medium) AND difficult fixtures(low) could be more serious severity potentially!
		switch signal.Type {
		case SignalAvailabilityRisk:
			if signal.Severity == SeverityHigh {
				if available[player.ElementType] < minimum[player.ElementType] {
					priority = SeverityHigh
					reason = "Squad cannot field the minimum number of available players in this position. " + signal.Reason
				} else {
					priority = SeverityMedium
					reason = "Player unavailable, but squad has sufficient positional cover. " + signal.Reason
				}
			} else {
				priority = SeverityLow
				reason = "Player unavailability uncertain. " + signal.Reason
			}
		case SignalPoorForm:
			if signal.Severity == SeverityMedium {
				priority = SeverityMedium
				reason = "Player's recent form warrants reviewing their place in the squad. " + signal.Reason
			} else {
				continue
			}
		case SignalDifficultFixtures:
			priority = SeverityLow
			reason = "Player has a difficult upcoming fixture run. " + signal.Reason
		default:
			continue
		}

		issues = append(issues, SquadIssue{
			PlayerID: player.ID,
			Priority: priority,
			Reason: reason,
		})
	}

	return issues
}


// AnalyseManagerSquad acts as an orchestrator in calculating fixture difficulty for the squad, evaluating signals
// and return list of potential squad issues used for recommendation engine
func AnalyseManagerSquad(
	squad SquadDetails,
	fixtures []Fixture,
	nextGameweek int,
	gameweeks int,
) []SquadIssue {
	fixtureAnalysis := AnalyseSquadFixtures(
		squad,
		fixtures,
		nextGameweek,
		gameweeks,
		)

	signals := EvaluateSquad(squad, fixtureAnalysis)

	issues := IdentifySquadIssues(squad, signals)

	return issues
}
