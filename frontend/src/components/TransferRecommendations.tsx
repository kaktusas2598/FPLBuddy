import { useEffect, useState } from 'react';
import type { TransferRecommendation } from '../types.ts';

interface Props {
    managerId: number;
    gameweekId: number;
}

function TransferRecommendations({ managerId, gameweekId }: Props) {
    const [recommendations, setRecommendations] = useState<TransferRecommendation[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        const controller = new AbortController();

        async function fetchRecommendations() {
            setLoading(true);
            setError(null);
            setRecommendations([]);

            const startGameweek = gameweekId + 1;
            const gameweekCount = 5;

            try {
                const response = await fetch(
                    `/api/manager/${managerId}/transfer-recommendations?gameweek=${gameweekId}&start=${startGameweek}&count=${gameweekCount}`,
                    { signal: controller.signal }
                );

                if (!response.ok) {
                    throw new Error(`Transfer recommendations lookup failed: ${response.status}`);
                }

                const result: TransferRecommendation[] = await response.json();
                if (!controller.signal.aborted) {
                    setRecommendations(result ?? []);
                }
            } catch (err) {
                if (!controller.signal.aborted) {
                    setError(err instanceof Error ? err.message : 'Failed to load recommendations');
                }
            } finally {
                if (!controller.signal.aborted) {
                    setLoading(false);
                }
            }
        }

        fetchRecommendations();

        return () => controller.abort();
    }, [managerId, gameweekId]);

    return (
        <section className="manager-panel transfer-recommendations">
            <h2>Transfer Recommendations</h2>
            <p>Potential transfers for GW{gameweekId + 1}–{gameweekId + 5}</p>

            {loading && <p>Analysing your squad...</p>}
            {error && <p>{error}</p>}

            {!loading && !error && recommendations.length === 0 && (
                <p>No transfers currently recommended.</p>
            )}

            {!loading && !error && (
                <div className="transfer-recommendations-grid">
                    {recommendations.map((recommendation) => (
                        <div className="transfer-card" key={recommendation.player_out.id}>
                            <h3>{recommendation.player_out.web_name}</h3>

                            <div className="transfer-issues">
                                {recommendation.issues.map((issue, index) => (
                                    <p key={index}>
                                        <strong>{issue.priority}:</strong> {issue.reason}
                                    </p>
                                ))}
                            </div>

                            <h4>Potential replacements</h4>

                            <div className="performance-grid">
                                {recommendation.replacements.slice(0, 3).map((player) => (
                                    <div key={player.id}>
                                        <span>{player.web_name}</span>
                                        <strong>£{(player.now_cost / 10).toFixed(1)}m</strong>
                                        <span>Form: {player.form}</span>
                                    </div>
                                ))}
                            </div>
                        </div>
                    ))}
                </div>
            )}

            <p className="transfer-disclaimer">
                Recommendations use current market prices rather than your actual selling prices.
                Some suggested transfers may exceed your available budget.
            </p>
        </section>
    );
}

export default TransferRecommendations;
