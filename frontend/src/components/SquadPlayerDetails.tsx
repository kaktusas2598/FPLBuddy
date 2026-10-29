import type { Squad } from '../types.ts';

type SquadPick = Squad['picks'][number];

function SquadPlayerDetails({ selectedPlayer, fixtureDifficulty }: {
    selectedPlayer: SquadPick;
    fixtureDifficulty: number;
}) {

    return (
        <div className="squad-player-details">
            <h2>{selectedPlayer.player.web_name}</h2>
            <p>Next 5 fixture average difficulty: {fixtureDifficulty}</p>
        </div>
    );
}

export default SquadPlayerDetails
