// React hooks
import { useState, useEffect } from 'react'
import premierLeagueLogo from './assets/premier-league-logo.png'
import './App.css'

// Custom components
import ManagerLookup from './components/ManagerLookup.tsx'
import CurrentGameweek from './components/CurrentGameweek.tsx'

interface Team {
    id: number;
    code: number;
    name: string;
    short_name: string;
}

// App() is a React component
function App() {
  // Creates status variable and setStatus func to update it and tell React to render it
  const [status, setStatus] = useState<string>("Checking..")

  const [teams, setTeams] = useState<Team[]>([]);

  useEffect(() => {
      async function checkServerHealth() {
          const healthResponse = await fetch('/api/health');
          const healthResult = await healthResponse.text();

          setStatus(healthResult)
      }

      async function getTeams() {
          const teamsResponse = await fetch('/api/teams');
          const teamsResult = await teamsResponse.json();

          setTeams(teamsResult)
      }

      checkServerHealth()
      getTeams()
  }, []);

  // This HTML like syntax is JSX or for typescript its TSX
  return (
    <>
      <main className="app-main">
          <section id="center">
            <div className="hero">
              <img src={premierLeagueLogo} className="base" width="170" height="179" alt="" />
            </div>
            <div>
              <h1>FPL Buddy</h1>
              <p>Fantasy Premier League Assistant</p>
          </div>
          </section>

          <CurrentGameweek />
          <ManagerLookup />

          <p>Premier League Teams</p>
          <div className="teams-grid">
            {teams.map((team) => (
              <div className="team-card" key={team.id}>
                <img
                    src={`https://resources.premierleague.com/premierleague25/badges/${team.code}.svg`}
                    alt={`${team.name} badge`}
                    width="48"
                    height="48"
                />
                <h3>{team.name}</h3>
                <span>{team.short_name}</span>
              </div>
            ))}
          </div>

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
