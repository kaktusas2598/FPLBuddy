# FPLBuddy

**FPLBuddy** is a work-in-progress Fantasy Premier League assistant app that helps managers understand their squads, assess upcoming fixtures and, eventually, make better transfer and team-selection decisions.

Built as a personal project to help me learn **Go**, **React** and **TypeScript**, using data from the official Fantasy Premier League API.

## Features

- Look up an FPL manager and view their squad on a pitch.
- Inspect player performance, availability and position-specific statistics.
- Review upcoming fixtures, opponents and average fixture difficulty.
- Compare squad metrics including points, form, price, expected goals and ownership.
- Fetch current gameweek, team and player data through a Go API.
- Cache FPL bootstrap data in memory to reduce repeated API requests.

## Tech stack

- **Backend:** Go, standard library HTTP server
- **Frontend:** React, TypeScript, Vite
- **Development:** Docker Compose
- **Data:** Official Fantasy Premier League API
- **Tools** NeoVim

## Roadmap

- [ ] Position-specific player rankings
- [ ] Squad analysis and keep/sell/monitor recommendations
- [ ] Transfer and replacement suggestions
- [ ] Starting XI and captaincy recommendations
- [ ] Historical data and predictive modelling

## Project status

FPLBuddy is under active development. The current focus is turning player and fixture data into useful, explainable recommendations.

## Disclaimer

FPLBuddy is an independent project and is not affiliated with or endorsed by the Premier League or Fantasy Premier League.

