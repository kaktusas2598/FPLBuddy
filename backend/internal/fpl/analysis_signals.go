package fpl

// EvaluatePlayerAvailability determines if provided player is available for a match and returns a new PlayerSignal
func EvaluatePlayerAvailability(player Player) []PlayerSignal {

	var playerSignals []PlayerSignal
	reason := player.News

	// NOTE: In go break statements are not needed to avoid fall through to the next case!
	switch player.Status {
	case "a": // Available
		return nil
	case "d": // Doubtful
		if reason == "" {
			reason = "Player is currently doubtful"
		}

		playerSignals = append(playerSignals, PlayerSignal{
			PlayerID: player.ID,
			Type:     SignalAvailabilityRisk,
			Severity: SeverityMedium,
			Reason:   reason,
		})
	case "i": // Injured
	case "s": // Suspended
	case "u": // Unavailable
		if reason == "" {
			reason = "Player is currently injured, suspended or unavailable"
		}

		playerSignals = append(playerSignals, PlayerSignal{
			PlayerID: player.ID,
			Type:     SignalAvailabilityRisk,
			Severity: SeverityHigh,
			Reason:   reason,
		})
	default:
	}

	return playerSignals
}
