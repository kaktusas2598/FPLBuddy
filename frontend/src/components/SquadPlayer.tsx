import type { Squad } from '../types.ts';
import type { SquadMetric } from '../types.ts';

// Custom type for this component only
type SquadPick = Squad['picks'][number];
// Custom type + extra props passed into this component
type SquadPlayerProps = SquadPick & {
    selectedMetric: SquadMetric;
}

function SquadPlayer({ player, pick, selectedMetric }: SquadPlayerProps) {

    let detailsValue: string | number;
    switch (selectedMetric) {
        case 'points':
            detailsValue = player.event_points * (pick.multiplier || 1);
            break;
        case 'price':
            detailsValue = `£${(player.now_cost / 10).toFixed(1)}m`;
            break;
        case 'price_change':
            detailsValue = `${player.cost_change_start > 0 ? '+' : ''}£${(player.cost_change_start / 10).toFixed(1)}m`;
            break;
        case 'form':
            detailsValue = player.form;
            break;
        case 'xg':
            detailsValue = player.expected_goals;
            break;
        case 'ownership':
            detailsValue = player.selected_by_percent + '%';
            break;
        default:
            detailsValue = 'N/A';
    }

    return (
        <div className="squad-player">
            <img
                className="squad-player-badge"
                src={`https://resources.premierleague.com/premierleague25/badges/${player.team_code}.svg`}
                alt="Club badge"
            />

            <span>
                {player.web_name}
                {pick.is_captain && ' (C)'}
                {pick.is_vice_captain && ' (VC)'}
            </span>
            <span className="squad-player-subdetails">
                {detailsValue}
            </span>
        </div>
    );
}

export default SquadPlayer
