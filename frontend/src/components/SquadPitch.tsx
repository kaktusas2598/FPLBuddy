import type { Squad } from '../types.ts';
import SquadPlayer from './SquadPlayer';

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
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                        />
                    ))}
                </div>

                <div className="squad-row">
                    {defenders.map(({ player, pick }) => (
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                        />
                    ))}
                </div>
                <div className="squad-row">
                    {midfielders.map(({ player, pick }) => (
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                        />
                    ))}
                </div>
                <div className="squad-row">
                    {forwards.map(({ player, pick }) => (
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                        />
                    ))}
                </div>
            </div>
            <div className="squad-bench">
                <div className="squad-row">
                    {substitutes.map(({ player, pick }) => (
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                        />
                    ))}
                </div>

            </div>
        </>
    )
}

export default SquadPitch

