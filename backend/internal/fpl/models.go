package fpl

// Player also known as "Element" via official FPL API, but I like Player better!
type Player struct {
	// Identity
	ID          int    `json:"id"`
	Code        int    `json:"code"`
	FirstName   string `json:"first_name"`
	SecondName  string `json:"second_name"`
	WebName     string `json:"web_name"`
	Team        int    `json:"team"`
	TeamCode    int    `json:"team_code"`
	ElementType int    `json:"element_type"`
	Photo       string `json:"photo"`

	// Price (values in tenths of £1m)
	NowCost             int `json:"now_cost"`
	CostChangeEvent     int `json:"cost_change_event"`
	CostChangeEventFall int `json:"cost_change_event_fall"`
	CostChangeStart     int `json:"cost_change_start"`
	CostChangeStartFall int `json:"cost_change_start_fall"`

	// Points and form
	TotalPoints   int    `json:"total_points"`
	EventPoints   int    `json:"event_points"`
	PointsPerGame string `json:"points_per_game"`
	Form          string `json:"form"`
	EPNext        string `json:"ep_next"`
	ValueForm     string `json:"value_form"`
	ValueSeason   string `json:"value_season"`

	// Ownership and transfers
	SelectedByPercent string `json:"selected_by_percent"`
	TransfersIn       int    `json:"transfers_in"`
	TransfersOut      int    `json:"transfers_out"`
	TransfersInEvent  int    `json:"transfers_in_event"`
	TransfersOutEvent int    `json:"transfers_out_event"`

	// Availability
	Status                   string `json:"status"`
	News                     string `json:"news"`
	ChanceOfPlayingNextRound *int   `json:"chance_of_playing_next_round"`
	ChanceOfPlayingThisRound *int   `json:"chance_of_playing_this_round"`

	// Playing time
	Minutes int `json:"minutes"`
	Starts  int `json:"starts"`

	// Attacking statistics
	GoalsScored    int  `json:"goals_scored"`
	Assists        int  `json:"assists"`
	PenaltiesOrder *int `json:"penalties_order"`

	// Expected attacking statistics
	ExpectedGoals            string  `json:"expected_goals"`
	ExpectedAssists          string  `json:"expected_assists"`
	ExpectedGoalInvolvements string  `json:"expected_goal_involvements"`
	ExpectedGoalsPer90       float64 `json:"expected_goals_per_90"`
	ExpectedAssistsPer90     float64 `json:"expected_assists_per_90"`

	// Defensive statistics
	CleanSheets           int `json:"clean_sheets"`
	GoalsConceded         int `json:"goals_conceded"`
	OwnGoals              int `json:"own_goals"`
	Saves                 int `json:"saves"`
	PenaltiesSaved        int `json:"penalties_saved"`
	PenaltiesMissed       int `json:"penalties_missed"`
	DefensiveContribution int `json:"defensive_contribution"`

	// Expected defensive statistics
	ExpectedGoalsConceded      string  `json:"expected_goals_conceded"`
	ExpectedGoalsConcededPer90 float64 `json:"expected_goals_conceded_per_90"`
	CleanSheetsPer90           float64 `json:"clean_sheets_per_90"`

	// Discipline
	YellowCards int `json:"yellow_cards"`
	RedCards    int `json:"red_cards"`

	// Bonus points
	Bonus          int  `json:"bonus"`
	BPS            int  `json:"bps"`
	DreamteamCount int  `json:"dreamteam_count"`
	InDreamteam    bool `json:"in_dreamteam"`

	// ICT metrics
	Influence  string `json:"influence"`
	Creativity string `json:"creativity"`
	Threat     string `json:"threat"`
	ICTIndex   string `json:"ict_index"`
}

// Fixture represents a single football match for a current season
// TODO: potentially want a different type for "Fixture history" extracted from /api/element-summary/{player-id}
type Fixture struct {
	ID                  int    `json:"id"`
	Code                int    `json:"code"`
	GameWeek            int    `json:"event"`
	Finished            bool   `json:"finished"`
	FinishedProvisional bool   `json:"finished_provisional"`
	KickoffTime         string `json:"kickoff_time"`
	Minutes             int    `json:"minutes"`
	Started             bool   `json:"started"`
	TeamA               int    `json:"team_a"`
	TeamAScore          int    `json:"team_a_score"`
	TeamADifficulty     int    `json:"team_a_difficulty"`
	TeamH               int    `json:"team_h"`
	TeamHScore          int    `json:"team_h_score"`
	TeamHDifficulty     int    `json:"team_h_difficulty"`
	// TODO: Potentially want to add `stats` as well if needed for later
}

// Team represents actual premier league team, not to be confused with FPL squad!
type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	Code      int    `json:"code"` // Matches to SVG path from official FPL API to extract the logo
}

// GameWeek or "Event" represents a single gameweek in FPL in FPLBuddy internal API
type GameWeek struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DeadlineTime string `json:"deadline_time"`
	MostPoints   int    `json:"top_element"`
	MostSelected int    `json:"most_selected"`
	IsCurrent    bool   `json:"is_current"`
}

// GameWeekDetails represents single gameweek in FPL with enriched details for presentation layer in FPLBuddy
type GameWeekDetails struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DeadlineTime string `json:"deadline_time"`
	MostPoints   Player `json:"top_element"`
	MostSelected Player `json:"most_selected"`
	IsCurrent    bool   `json:"is_current"`
}

// Pick represents prem player picked by a manager for their squad
type Pick struct {
	PlayerID      int  `json:"element"`
	Position      int  `json:"position"`
	Multiplier    int  `json:"multiplier"` // 0 if benched
	IsCaptain     bool `json:"is_captain"`
	IsViceCaptain bool `json:"is_vice_captain"`
}

// SquadPlayer combines player/element information with manager's gameweek pick
type SquadPlayer struct {
	Player Player `json:"player"`
	Pick   Pick   `json:"pick"`
}

// Squad represents selection of 15 picks/players chosen by FPL manager for a gameweek
type Squad struct {
	Picks []Pick `json:"picks"`
}

// SquadDetails containes enriched information about FPL manager squad selection including players details
type SquadDetails struct {
	Picks []SquadPlayer `json:"picks"`
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
