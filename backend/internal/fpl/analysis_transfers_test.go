package fpl

import "testing"

func TestFindReplacementCandidates(t *testing.T) {
	playerOut := Player{
		ID:          1,
		ElementType: 3,
		Team:        1,
		NowCost:     75,
		Status:      "a",
	}

	squad := SquadDetails{
		Picks: []SquadPlayer{
			{Player: playerOut},
			{Player: Player{ID: 2, ElementType: 3, Team: 2}},
			{Player: Player{ID: 3, ElementType: 2, Team: 2}},
			{Player: Player{ID: 4, ElementType: 4, Team: 2}},
			{Player: Player{ID: 5, ElementType: 2, Team: 1}},
			{Player: Player{ID: 6, ElementType: 1, Team: 1}},
		},
	}

	tests := []struct {
		name         string
		candidate    Player
		bank         int
		sellingPrice int
		wantIncluded bool
	}{
		{
			name: "eligible affordable midfielder",
			candidate: Player{
				ID: 10, ElementType: 3, Team: 3,
				Status: "a", NowCost: 85,
			},
			bank: 20, sellingPrice: 70,
			wantIncluded: true,
		},
		{
			name: "wrong position",
			candidate: Player{
				ID: 11, ElementType: 4, Team: 3,
				Status: "a", NowCost: 80,
			},
			bank: 20, sellingPrice: 70,
			wantIncluded: false,
		},
		{
			name: "already in squad",
			candidate: Player{
				ID: 2, ElementType: 3, Team: 2,
				Status: "a", NowCost: 60,
			},
			bank: 20, sellingPrice: 70,
			wantIncluded: false,
		},
		{
			name: "unavailable player",
			candidate: Player{
				ID: 12, ElementType: 3, Team: 3,
				Status: "i", NowCost: 70,
			},
			bank: 20, sellingPrice: 70,
			wantIncluded: false,
		},
		{
			name: "club already has three players",
			candidate: Player{
				ID: 13, ElementType: 3, Team: 2,
				Status: "a", NowCost: 70,
			},
			bank: 20, sellingPrice: 70,
			wantIncluded: false,
		},
		{
			name: "outgoing player's club remains eligible",
			candidate: Player{
				ID: 14, ElementType: 3, Team: 1,
				Status: "a", NowCost: 70,
			},
			bank: 20, sellingPrice: 70,
			wantIncluded: true,
		},
		{
			name: "exactly affordable",
			candidate: Player{
				ID: 15, ElementType: 3, Team: 3,
				Status: "a", NowCost: 90,
			},
			bank: 20, sellingPrice: 70,
			wantIncluded: true,
		},
		{
			name: "exceeds budget",
			candidate: Player{
				ID: 16, ElementType: 3, Team: 3,
				Status: "a", NowCost: 91,
			},
			bank: 20, sellingPrice: 70,
			wantIncluded: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			candidates := FindReplacementCandidates(
				playerOut,
				squad,
				[]Player{tt.candidate},
				tt.bank,
				tt.sellingPrice,
			)

			gotIncluded := len(candidates) == 1

			if gotIncluded != tt.wantIncluded {
				t.Errorf(
					"candidate %d: included = %v, want %v",
					tt.candidate.ID,
					gotIncluded,
					tt.wantIncluded,
				)
			}
		})
	}
}
