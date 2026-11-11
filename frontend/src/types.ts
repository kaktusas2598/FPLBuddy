export interface Player {
    id: number;
    web_name: string;
    team: number;
    team_code: number;
    element_type: number;
    now_cost: number; // Divide by 10
    selected_by_percent: string;
    cost_change_start: number;
    // Performance statistics
    form: string;
    total_points: number;
    event_points: number; // To get per manager, account for chips used and captaincy
    points_per_game: string;
    expected_assists: string;
    expected_goals: string;
    goals_scored: number;
    assists: number;
    minutes: number;
    // Defensive statistics
    clean_sheets: number;
    goals_conceded: number;
    defensive_contribution: number;
    saves: number;
    // Bonus points
    bonus: number;
    // Injury/availability related information
    status: string;
    news: string;
    chance_of_playing_next_round: number | null;
    chance_of_playing_this_round: number | null;
}

export interface Squad {
    picks: {
        player: Player;
        pick: {
            element: number;
            position: number;
            multiplier: number;
            is_captain: boolean;
            is_vice_captain: boolean;
        };
    }[];
}

export interface Fixture {
    id: number;
    code: number;
    event: number | null;
    finished: boolean;
    finished_provisional: boolean;
    kickoff_time: string;
    minutes: number;
    started: boolean;
    team_a: number;
    team_h: number;
    team_a_score: number | null;
    team_h_score: number | null;
    team_a_difficulty: number;
    team_h_difficulty: number;
}

export interface PlayerFixtureAnalysis {
    player_id: number;
    average_difficulty: number;
    fixtures: Fixture[];
}

export interface Team {
    id: number;
    name: string;
    short_name: string;
    code: number;
}

export type SignalSeverity = "low" | "medium" | "high";

export interface SquadIssue {
  player_id: number;
  priority: SignalSeverity;
  reason: string;
}

export interface TransferRecommendation {
  player_out: Player;
  replacements: Player[];
  issues: SquadIssue[];
}

export type SquadMetric = 'points' | 'price' | 'price_change' | 'form' | 'xg' | 'ownership'
