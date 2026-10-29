import type { Squad } from '../types.ts';

type SquadPick = Squad['picks'][number];

function SquadPlayerDetails({ selectedPlayer }: {
    selectedPlayer: SquadPick;
}) {

    return (
        <div className="squad-player-details">
            <h2>{selectedPlayer.player.web_name}</h2>
        </div>
    );
}

export default SquadPlayerDetails
