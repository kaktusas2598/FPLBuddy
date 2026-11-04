import type { Squad } from '../types.ts';
import type { Fixture } from '../types.ts';

type SquadPick = Squad['picks'][number];

function SquadPlayerDetails({ selectedPlayer, fixtureDifficulty, fixtures }: {
    selectedPlayer: SquadPick;
    fixtureDifficulty: number;
    fixtures: Fixture[];
}) {

    return (
        <div className="squad-player-details">
            <h2>{selectedPlayer.player.web_name}</h2>
            <p>Next 5 fixture average difficulty: {fixtureDifficulty}</p>
            <div className="fixture-list">
                {fixtures.map((fixture) => {
                    const isHome = fixture.team_h === selectedPlayer.player.team;

                    const difficulty = isHome
                        ? fixture.team_h_difficulty
                        : fixture.team_a_difficulty;

                    return (
                        <div key={fixture.id} className={`fixture-badge fixture-difficulty-${difficulty}`}>
                            GW{fixture.event}: {difficulty}
                        </div>
                    );
                })}
            </div>

        </div>
    );
}

export default SquadPlayerDetails
