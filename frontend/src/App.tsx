// React hooks
import { useState, useEffect } from 'react'
import heroImg from './assets/hero.png'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import premierLeagueLogo from './assets/premier-league-logo.png'
import './App.css'

interface Team {
    id: number;
    name: string;
    short_name: string;
}

// App() is a React component
function App() {
  const [count, setCount] = useState(0)

  // Creates status variable and setStatus func to update it and tell React to render it
  const [status, setStatus] = useState<string>("Checking..")


  const [players, setPlayers] = useState<string>("");
  const [teams, setTeams] = useState<Team[]>([]);

  useEffect(() => {
      async function checkServerHealth() {
          const healthResponse = await fetch('/api/health');
          const healthResult = await healthResponse.text();

          setStatus(healthResult)
      }

      async function getPlayers() {
          const playersResponse = await fetch('/api/players');
          const playersResult = await playersResponse.text();

          setPlayers(playersResult)
      }

      async function getTeams() {
          const teamsResponse = await fetch('/api/teams');
          const teamsResult = await teamsResponse.json();

          setTeams(teamsResult)
      }



      checkServerHealth()
      getPlayers()
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
            <button
              type="button"
              className="counter"
              onClick={() => setCount((count) => count + 1)}
            >
              Count is {count}
            </button>
          </section>

          <div className="ticks"></div>

          <div className="teams-grid">
            {teams.map((team) => (
              <div className="team-card" key={team.id}>
                <h3>{team.name}</h3>
                <span>{team.short_name}</span>
              </div>
            ))}
          </div>

          <div className="ticks"></div>

          <section id="next-steps">
            <div id="docs">
              <svg className="icon" role="presentation" aria-hidden="true">
                <use href="/icons.svg#documentation-icon"></use>
              </svg>
              <h2>Documentation</h2>
              <ul>
                <li>
                  <a href="https://vite.dev/" target="_blank">
                    <img className="logo" src={viteLogo} alt="" />
                    Explore Vite
                  </a>
                </li>
                <li>
                  <a href="https://react.dev/" target="_blank">
                    <img className="button-icon" src={reactLogo} alt="" />
                    Learn more
                  </a>
                </li>
                <li>
                  <a href="https://go.dev/doc/tutorial/" target="_blank">
                    GoLang Tutorials
                  </a>
                </li>
              </ul>
            </div>
          </section>

          <div className="ticks"></div>
          <section id="spacer"></section>
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
