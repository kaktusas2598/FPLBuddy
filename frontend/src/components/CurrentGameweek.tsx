import { useState, useEffect } from 'react'

interface GameWeek {
    id: number,
    name: string;
    most_selected: { web_name: string};
    top_element: { web_name: string};
    deadline_time: string;
}

function CurrentGameweek() {
    const [currentGameweek, setCurrentGameweek] = useState<GameWeek | null>(null);

    useEffect(() => {
        async function fetchCurrentGameweek() {
            const currentGameweekResponse = await fetch('/api/current-gameweek');
            if (!currentGameweekResponse.ok) {
                throw new Error(`Current GameWeek lookup failed: ${currentGameweekResponse.status}`);
            }

            const currentGameweekResult = await currentGameweekResponse.json();
            setCurrentGameweek(currentGameweekResult)
            console.log(currentGameweekResult);
        }

        fetchCurrentGameweek();

    }, []);

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
