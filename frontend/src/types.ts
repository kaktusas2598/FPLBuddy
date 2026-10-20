export interface Squad {
    picks: {
        player: {
            id: number;
            web_name: string;
            team: number;
            team_code: number;
            element_type: number;
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

