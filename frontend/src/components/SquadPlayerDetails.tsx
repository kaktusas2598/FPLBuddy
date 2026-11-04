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

    const statusLabels: Record<string, string> = {
        a: 'Available',
        d: 'Doubtful',
        i: 'Injured',
        s: 'Suspended',
        u: 'Unavailable',
    };

    return (
        <div className="squad-player-details">
            <h2>Selected Player: {selectedPlayer.player.web_name}</h2>

            {selectedPlayer.player.status !== 'a' && (
                <div className="player-availability">
                    <h3>Availability</h3>
                    <strong>{statusLabels[selectedPlayer.player.status] ?? 'Unavailable'}</strong>

                    {selectedPlayer.player.news && (
                        <p>{selectedPlayer.player.news}</p>
                    )}

                    {selectedPlayer.player.chance_of_playing_next_round !== null && (
                        <span>
                            Chance of playing next round: {selectedPlayer.player.chance_of_playing_next_round}%
                        </span>
                    )}
                </div>
            )}

            {/* TODO: Show position-specific performance statistics, for example saves for GK and Defensive contributions for DEF */}
            {/* TODO: Show player rank for each indidual statistics, for ex. "Total points (3rd): 123" */}
            <div className="player-performance">
                <h3>Performance</h3>
                <div className="performance-grid">
                    <div>Total points <strong>{selectedPlayer.player.total_points}</strong></div>
                    <div>Points per game <strong>{selectedPlayer.player.points_per_game}</strong></div>
                    <div>Form <strong>{selectedPlayer.player.form}</strong></div>
                    <div>Minutes <strong>{selectedPlayer.player.minutes}</strong></div>
                    <div>Goals <strong>{selectedPlayer.player.goals_scored}</strong></div>
                    <div>Assists <strong>{selectedPlayer.player.assists}</strong></div>
                    <div>xG <strong>{selectedPlayer.player.expected_goals}</strong></div>
                    <div>xA <strong>{selectedPlayer.player.expected_assists}</strong></div>
                </div>
            </div>

            <div className="player-fixtures">
                <h3>Upcoming Fixtures</h3>
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

        </div>
    );
}

export default SquadPlayerDetails
