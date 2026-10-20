import { useState } from 'react'

import type { Squad } from '../types.ts';
import type { SquadMetric } from '../types.ts';
import SquadPlayer from './SquadPlayer';


function SquadPitch({ squad }: { squad: Squad}) {
    const [selectedMetric, setSelectedMetric] = useState<SquadMetric>('points');

    function handleMetricChange(e: React.ChangeEvent<HTMLSelectElement>) {
        e.preventDefault();

        setSelectedMetric(e.target.value as SquadMetric);
    }

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
                <select className="squad-metric-select" onChange={handleMetricChange} value={selectedMetric}>
                    <option value="points">Gameweek Points</option>
                    <option value="price">Current Price</option>
                    <option value="price_change">Season Price Change</option> {/* TODO: also want gw price change */}
                    <option value="form">Form</option>
                    <option value="ownership">Ownership</option>
                    <option value="xg">Season xG</option> {/* TODO: probably better to display xG per 90min? */}
                </select>
                <div className="squad-row">
                    {goalkeepers.map(({ player, pick }) => (
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                            selectedMetric={selectedMetric}
                        />
                    ))}
                </div>

                <div className="squad-row">
                    {defenders.map(({ player, pick }) => (
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                            selectedMetric={selectedMetric}
                        />
                    ))}
                </div>
                <div className="squad-row">
                    {midfielders.map(({ player, pick }) => (
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                            selectedMetric={selectedMetric}
                        />
                    ))}
                </div>
                <div className="squad-row">
                    {forwards.map(({ player, pick }) => (
                        <SquadPlayer
                            key={player.id}
                            player={player}
                            pick={pick}
                            selectedMetric={selectedMetric}
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
                            selectedMetric={selectedMetric}
                        />
                    ))}
                </div>

            </div>
        </>
    )
}

export default SquadPitch

