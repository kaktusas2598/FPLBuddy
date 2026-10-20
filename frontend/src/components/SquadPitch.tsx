import type { Squad } from '../types.ts';

function SquadPitch({ squad }: { squad: Squad}) {

    // Separate picks by starting/non-starting and by their role on the pitch
    const starters = squad.picks.filter(({ pick }) => pick.position <= 11);
    const substitutes = squad.picks.filter(({ pick }) => pick.position > 11);

    const goalkeepers = starters.filter(({ player }) => player.element_type === 1);
    const defenders = starters.filter(({ player }) => player.element_type === 2);
    const midfielders = starters.filter(({ player }) => player.element_type === 3);
    const forwards = starters.filter(({ player }) => player.element_type === 4);

    return (
        <>
            <div className="squad-pitch">
                <div className="squad-row">
                    {goalkeepers.map(({ player, pick }) => (
                        <div className="squad-player" key={player.id}>
                            {player.web_name}
                            {pick.is_captain && ' (C)'}
                            {pick.is_vice_captain && ' (VC)'}
                        </div>
                    ))}
                </div>

                <div className="squad-row">
                    {defenders.map(({ player, pick }) => (
                        <div className="squad-player" key={player.id}>
                            {player.web_name}
                            {pick.is_captain && ' (C)'}
                            {pick.is_vice_captain && ' (VC)'}
                        </div>
                    ))}
                </div>
                <div className="squad-row">
                    {midfielders.map(({ player, pick }) => (
                        <div className="squad-player" key={player.id}>
                            {player.web_name}
                            {pick.is_captain && ' (C)'}
                            {pick.is_vice_captain && ' (VC)'}
                        </div>
                    ))}
                </div>
                <div className="squad-row">
                    {forwards.map(({ player, pick }) => (
                        <div className="squad-player" key={player.id}>
                            {player.web_name}
                            {pick.is_captain && ' (C)'}
                            {pick.is_vice_captain && ' (VC)'}
                        </div>
                    ))}
                </div>
            </div>
            <div className="squad-bench">
                <div className="squad-row">
                    {substitutes.map(({ player }) => (
                        <div className="squad-player" key={player.id}>
                            {player.web_name}
                        </div>
                    ))}
                </div>

            </div>
        </>
    )
}

export default SquadPitch

