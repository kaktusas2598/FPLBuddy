// React hook
import { useState } from 'react'
import heroImg from './assets/hero.png'
import reactLogo from './assets/react.svg'
import viteLogo from './assets/vite.svg'
import premierLeagueLogo from './assets/premier-league-logo.png'
import './App.css'

// App() is a React component
function App() {
  const [count, setCount] = useState(0)



  // This HTML like syntax is JSX or for typescript its TSX
  return (
    <>
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
          </ul>
        </div>
      </section>

      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  )
}

export default App
