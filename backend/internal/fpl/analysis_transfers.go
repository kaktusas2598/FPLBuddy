package fpl

// FindReplacementCandidates provided the player, finds the potential candidates to replace/transfer him
// function filters potential candidates but is not responsible for ranking them
func FindReplacementCandidates(
	playerOut Player,
	squad SquadDetails,
	players []Player,
	bank int,
	sellingPrice int,
) []Player {

	var candidates []Player

	// Maps for player IDs in the squad and number of players for each club
	squadPlayerIDs := make(map[int]bool)
	clubCounts := make(map[int]int)

	for _, pick := range squad.Picks {
		squadPlayerIDs[pick.Player.ID] = true
		// Do not count player out towards club count
		if pick.Player.ID != playerOut.ID {
			clubCounts[pick.Player.Team]++
		}
	}

	maxBudget := bank+sellingPrice

	for _, player := range players {
		if player.ElementType == playerOut.ElementType &&
			player.ID != playerOut.ID &&
			player.Status == "a" &&
			!squadPlayerIDs[player.ID] &&
			clubCounts[player.Team] < 3 &&
			player.NowCost <= maxBudget {
			candidates = append(candidates, player)
		}
	}

	return candidates
}
