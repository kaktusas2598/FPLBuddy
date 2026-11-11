package fpl

import (
	"sort"
	"strconv"
)

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

// RankReplacementCandidates ranks eligible transfer candidates using recent form and upcoming fixture difficulty.
// Returns candidates ordered from most to least promising.
func RankReplacementCandidates(
    candidates []Player,
    fixtures []Fixture,
    nextGameweek int,
    gameweeks int,
) []Player {


	// Custom local struct which combines each candidate and their calculated ranking score based on form and fixtures
	type scoredPlayer struct {
		player Player
		score  float64
	}

	var scoredCandidates []scoredPlayer

	// Calculate provisional score for each candidate based on recent form and upcoming fixture difficulty
	for _, candidate := range candidates {
		form, err := strconv.ParseFloat(candidate.Form, 64)
		if err != nil {
			continue
		}

		difficulty, matches := AnalyseFixtureDifficulty(candidate.Team, fixtures, nextGameweek, gameweeks)
		if len(matches) == 0 {
			continue
		}

		// NOTE: Provisional scoring formula!!
		score := form - difficulty

		scoredCandidates = append(scoredCandidates, scoredPlayer{
			player: candidate,
			score: score,
		})
	}

	sort.Slice(scoredCandidates, func(i, j int) bool {
		return scoredCandidates[i].score > scoredCandidates[j].score
	})


	// Preallocate slice for ranked candidate Players based on the size of scored candidates
	rankedCandidates := make([]Player, 0, len(scoredCandidates))
	for _, candidate := range scoredCandidates {
		rankedCandidates = append(rankedCandidates, candidate.player)
	}

    return rankedCandidates
}
