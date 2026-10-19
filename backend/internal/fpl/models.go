package fpl

// Player also known as "Element" via official FPL API, but I like Player better!
type Player struct {
	ID      int    `json:"id"`
	WebName string `json:"web_name"`
	Team    int    `json:"team"`
}

// Team represents actual premier league team, not to be confused with FPL squad!
type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Code      int    `json:"code"` // Matches to SVG path from official FPL API to extract the logo
}

// GameWeek TODO: this is currently based on GraphQL response, so might need changing
type GameWeek struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DeadlineTime string `json:"deadline_time"`
	MostPoints   Player `json:"top_element"`
	MostSelected Player `json:"most_selected"`
}

// Pick represents prem player picked by a manager for their squad
type Pick struct {
	PlayerID      int  `json:"element"`
	Position      int  `json:"position"`
	Multiplier    int  `json:"multiplier"` // 0 if benched
	IsCaptain     bool `json:"is_captain"`
	IsViceCaptain bool `json:"is_vice_captain"`
}

// Squad represents selection of 15 picks/players chosen by FPL manager for a gameweek
type Squad struct {
	Picks []Pick `json:"picks"`
}

// Manager represents individual managing FPL squad
type Manager struct {
	ID           int    `json:"id"`
	FirstName    string `json:"player_first_name"`
	LastName     string `json:"player_last_name"`
	TeamName     string `json:"name"`
	FavTeamID    int    `json:"favourite_team"`
	ClubBadgeSrc string `json:"club_badge_src"`
	OverallPts   int    `json:"summary_overall_points"`
	GameweekPts  int    `json:"summary_event_points"`
	OverallRank  int    `json:"summary_overall_rank"`
	GameweekRank int    `json:"summary_event_rank"`
}
