import type { GameWeek } from '../App';

function CurrentGameweek({ currentGameweek }: {
    currentGameweek: GameWeek | null
}) {

    return (
        <>
            {currentGameweek && (
              <div className="manager-result">
                <h2>{currentGameweek.name}</h2>

                    <p>
                        Next deadline is: {new Date(currentGameweek.deadline_time).toLocaleString('en-GB', {
                            dateStyle: 'medium',
                            timeStyle: 'short',
                            timeZone: 'Europe/London'
                        })}
                    </p>

                <div className="manager-stats">
                  <div>
                    <span>Top Scoring Player</span>
                    <strong>{currentGameweek.top_element.web_name}</strong>
                  </div>

                  <div>
                    <span>Most Selected Player</span>
                    <strong>{currentGameweek.most_selected.web_name}</strong>
                  </div>
                </div>
              </div>
            )}
    </>
    )
}

export default CurrentGameweek;
