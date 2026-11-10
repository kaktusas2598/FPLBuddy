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
	case "i": // Injured
	case "s": // Suspended
	case "u": // Unavailable
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
