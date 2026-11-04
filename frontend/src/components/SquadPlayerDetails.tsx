import type { Squad } from '../types.ts';
import type { Fixture } from '../types.ts';
import type { Team } from '../types.ts';

type SquadPick = Squad['picks'][number];

function SquadPlayerDetails({ selectedPlayer, fixtureDifficulty, fixtures, teams }: {
    selectedPlayer: SquadPick;
    fixtureDifficulty: number;
    fixtures: Fixture[];
    teams: Team[];
}) {

    return (
        <div className="squad-player-details">
            <h2>{selectedPlayer.player.web_name}</h2>
            <p>Next 5 fixture average difficulty: <strong>{fixtureDifficulty}</strong></p>
            <div className="fixture-list">
                {fixtures.map((fixture) => {
                    const isHome = fixture.team_h === selectedPlayer.player.team;
                    const opponentId = isHome ? fixture.team_a : fixture.team_h;
                    const opponent = teams.find(team => team.id === opponentId);

                    const difficulty = isHome
                        ? fixture.team_h_difficulty
                        : fixture.team_a_difficulty;

                    return (
                        <div key={fixture.id} className={`fixture-badge fixture-difficulty-${difficulty}`}>
                            <span className="fixture-opponent">{opponent?.short_name ?? '???'}</span>
                            <span className="fixture-venue">({isHome ? 'H' : 'A'})</span>
                        </div>
                    );
                })}
            </div>

        </div>
    );
}

export default SquadPlayerDetails
