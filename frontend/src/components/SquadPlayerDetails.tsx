import type { Squad } from '../types.ts';
import type { Fixture } from '../types.ts';
import type { Team } from '../types.ts';
import type { Player } from '../types.ts';

type SquadPick = Squad['picks'][number];

// getRank counts how many players of the same position has higher value to calculate global rank for selected metric
// NOTE: potentially want to do this in Go backend depending on the speed? However for several hundred players
// so far this seems still very fast so not an immediate priority!
// NOTE: for points per game and form potentially consider players with few minutes ranking unusually highly?!
function getRank(players: Player[], selectedPlayer: Player, stat: keyof Player): number {
    return players.filter(
        p => p.element_type === selectedPlayer.element_type &&
             Number(p[stat]) > Number(selectedPlayer[stat])
    ).length + 1
}

function SquadPlayerDetails({ selectedPlayer, fixtureDifficulty, fixtures, teams, players }: {
    selectedPlayer: SquadPick;
    fixtureDifficulty: number;
    fixtures: Fixture[];
    teams: Team[];
    players: Player[];
}) {

    // NOTE: probably should move this to types.ts or sth like that as it wont change everytime component is loaded
    const statusLabels: Record<string, string> = {
        a: 'Available',
        d: 'Doubtful',
        i: 'Injured',
        s: 'Suspended',
        u: 'Unavailable',
    };

    // For convenience to make printing statistics shorter
    const player = selectedPlayer.player;
    // Also convenience short hand helper but also checks to make sure if players are loaded to avoid displaying '#1'
    const rank = (stat: keyof Player) =>
        players.length > 0 ? `#${getRank(players, player, stat)}` : '-';

    return (
        <div className="squad-player-details">
            <h2>Selected Player: {player.web_name}</h2>

            {player.status !== 'a' && (
                <div className="player-availability">
                    <h3>Availability</h3>
                    <strong>{statusLabels[player.status] ?? 'Unavailable'}</strong>

                    {player.news && (
                        <p>{player.news}</p>
                    )}

                    {player.chance_of_playing_next_round !== null && (
                        <span>
                            Chance of playing next round: {player.chance_of_playing_next_round}%
                        </span>
                    )}
                </div>
            )}

            <div className="player-performance">
                <h3>Performance</h3>
                <div className="performance-grid">
                    {/* Generic statistics sensible for all positions */}
                    <div>Total points <strong>{player.total_points} ({rank('total_points')})</strong></div>
                    <div>Points per game <strong>{player.points_per_game} ({rank('points_per_game')})</strong></div>
                    <div>Form <strong>{player.form} ({rank('form')})</strong></div>
                    <div>Minutes <strong>{player.minutes}</strong></div>
                    {/* Generic non-goalkeeper statistics */}
                    {player.element_type !== 1 && (
                        <>
                            <div>Goals <strong>{player.goals_scored} ({rank('goals_scored')})</strong></div>
                            <div>Assists <strong>{player.assists} ({rank('assists')})</strong></div>
                            <div>xG <strong>{player.expected_goals} ({rank('expected_goals')})</strong></div>
                            <div>xA <strong>{player.expected_assists} ({rank('expected_assists')})</strong></div>
                        </>
                    )}
                    {/* Goalkeeper-specific statistics */}
                    {player.element_type === 1 && (
                        <>
                            <div>Saves <strong>{player.saves} ({rank('saves')})</strong></div>
                            <div>Clean Sheets <strong>{player.clean_sheets} ({rank('clean_sheets')})</strong></div>
                            <div>Goals Conceded <strong>{player.goals_conceded}</strong></div>
                            <div>Bonus <strong>{player.bonus} ({rank('bonus')})</strong></div>
                        </>
                    )}
                    {/* Defender-specific statistics */}
                    {player.element_type === 2 && (
                        <>
                            <div>Clean Sheets <strong>{player.clean_sheets} ({rank('clean_sheets')})</strong></div>
                            <div>
                                Defensive contribution
                                <strong>{player.defensive_contribution} ({rank('defensive_contribution')})</strong>
                            </div>
                        </>
                    )}
                    {/* Midfielder-specific statistics */}
                    {player.element_type === 3 && (
                        <>
                            <div>
                                Defensive contribution
                                <strong>{player.defensive_contribution} ({rank('defensive_contribution')})</strong>
                            </div>
                        </>
                    )}

                </div>
            </div>

            <div className="player-fixtures">
                <h3>Upcoming Fixtures</h3>
                <p>Next 5 fixture average difficulty: <strong>{fixtureDifficulty}</strong></p>
                <div className="fixture-list">
                    {fixtures.map((fixture) => {
                        const isHome = fixture.team_h === player.team;
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
