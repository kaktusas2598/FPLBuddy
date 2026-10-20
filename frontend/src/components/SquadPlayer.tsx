import type { Squad } from '../types.ts';
type SquadPick = Squad['picks'][number];

function SquadPlayer({ player, pick }: SquadPick) {
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
        </div>
    );
}

export default SquadPlayer
