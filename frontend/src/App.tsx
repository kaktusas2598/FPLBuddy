// React hooks
import { useState, useEffect } from 'react'
import premierLeagueLogo from './assets/premier-league-logo.png'
import './App.css'

// Custom components
import ManagerLookup from './components/ManagerLookup.tsx'
import CurrentGameweek from './components/CurrentGameweek.tsx'
import SquadPitch from './components/SquadPitch.tsx'

// Custom types
import type { Squad } from './types.ts';

export interface GameWeek {
    id: number,
    name: string;
    most_selected: { web_name: string};
    top_element: { web_name: string};
    deadline_time: string;
}

// App() is a React component
function App() {
    const [status, setStatus] = useState<string>("Checking..")
    const [currentGameweek, setCurrentGameweek] = useState<GameWeek | null>(null);
    const [squad, setSquad] = useState<Squad | null>(null);

    useEffect(() => {
        async function checkServerHealth() {
            const healthResponse = await fetch('/api/health');
            const healthResult = await healthResponse.text();

            setStatus(healthResult)
        }

        async function fetchCurrentGameweek() {
            const currentGameweekResponse = await fetch('/api/current-gameweek');
            if (!currentGameweekResponse.ok) {
                throw new Error(`Current GameWeek lookup failed: ${currentGameweekResponse.status}`);
            }

            const currentGameweekResult = await currentGameweekResponse.json();
            setCurrentGameweek(currentGameweekResult)
            console.log(currentGameweekResult);
        }

        checkServerHealth()
        fetchCurrentGameweek();
    }, []);

    // This HTML like syntax is JSX or for typescript its TSX
    return (
        <>
            <main className="app-main">
                <header className="app-header">
                    <img src={premierLeagueLogo} alt="Premier League" />
                    <div>
                        <h1>FPL Buddy</h1>
                        <p>Fantasy Premier League Assistant</p>
                    </div>
                </header>

                <div className="dashboard-top">
                    <CurrentGameweek currentGameweek={currentGameweek}/>
                    <ManagerLookup gameweekId={currentGameweek?.id ?? null} setSquad={setSquad}/>
                </div>

                {squad && <SquadPitch squad={squad}/>}
            </main>

            <footer className="app-footer">
                <span>© {new Date().getFullYear()} Nerijus Vilčinskas</span>

                <span className="server-status">
                    Backend: <strong>{status}</strong>
                </span>
            </footer>
        </>
    )
}

export default App
