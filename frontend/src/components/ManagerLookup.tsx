import { useState } from 'react'

import SquadPitch from './SquadPitch.tsx'

interface Manager {
    name: string;
    player_first_name: string;
    player_last_name: string;
    summary_overall_points: number;
    summary_overall_rank: number;
}

export interface Squad {
    picks: {
        player: {
            id: number;
            web_name: string;
            team: number;
        };
        pick: {
            element: number;
            position: number;
            multiplier: number;
            is_captain: boolean;
            is_vice_captain: boolean;
        };
    }[];
}


function ManagerLookup({ gameweekId }: {
    gameweekId: number | null
}) {
    const [managerId, setManagerId] = useState('');
    const [manager, setManager] = useState<Manager | null>(null);
    const [squad, setSquad] = useState<Squad | null>(null);
    const [loading, setLoading] = useState(false);

    async function handleSubmit(e: React.SubmitEvent<HTMLFormElement>) {
        e.preventDefault();

        setLoading(true);
        setManager(null);
        setSquad(null);

        try {
            const managerResponse = await fetch('/api/manager/' + managerId);
            if (!managerResponse.ok) {
                throw new Error(`Manager lookup failed: ${managerResponse.status}`);
            }

            const managerResult = await managerResponse.json();
            setManager(managerResult)
            console.log(managerResult);

            if (gameweekId == null) {
                throw new Error('Current gameweek is not available');
            }
            const squadResponse = await fetch(`/api/manager/${managerId}/squad/${gameweekId}`);
            if (!squadResponse.ok) {
                throw new Error(`Squad lookup failed: ${squadResponse.status}`);
            }

            const squadResult = await squadResponse.json();
            setSquad(squadResult);
            console.log(squadResult);
        } catch (error) {
            console.log('Failed to fetch manager or squad:', error);
        } finally {
            setLoading(false);
        }
    }

    return (
        <>
            <form onSubmit={handleSubmit} className="manager-form">
              <label htmlFor="manager-id">Find your FPL team</label>

              <div className="manager-input-group">
                <input
                  id="manager-id"
                  type="number"
                  min="1"
                  placeholder="e.g. 1894731"
                  value={managerId}
                  onChange={(e) => setManagerId(e.target.value)}
                  required
                />

                <button type="submit" disabled={loading || gameweekId === null}>
                  {loading ? 'Searching...' : 'Find Team'}
                </button>
              </div>
            </form>

            {manager && (
              <div className="manager-result">
                <h2>{manager.name}</h2>

                <p>
                  Manager: {manager.player_first_name} {manager.player_last_name}
                </p>

                <div className="manager-stats">
                  <div>
                    <span>Overall points</span>
                    <strong>{manager.summary_overall_points.toLocaleString()}</strong>
                  </div>

                  <div>
                    <span>Overall rank</span>
                    <strong>{manager.summary_overall_rank.toLocaleString()}</strong>
                  </div>
                </div>
              </div>
            )}
            {squad && <SquadPitch squad={squad}/>}
        </>
    )
}

export default ManagerLookup;
