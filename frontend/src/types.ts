export interface Squad {
    picks: {
        player: {
            id: number;
            web_name: string;
            team: number;
            team_code: number;
            element_type: number;
            event_points: number; // To get per manager, account for chips used and captaincy
            now_cost: number; // Divide by 10
            form: string;
            selected_by_percent: string;
            cost_change_event: number;
            expected_goals: string;
        };
        pick: {
            element: number;
            position: number;
            multiplier: number;
            is_captain: boolean;
            is_vice_captain: boolean;
        };
    }[];
}

