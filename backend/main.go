package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/handlers"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	// FPL Gameweek Information
	http.HandleFunc("/current-gameweek", handlers.CurrentGameWeek)

	// FPL fixture information (all fixtures from the current season)
	http.HandleFunc("/fixtures", handlers.Fixtures)

	// FPL teams information
	http.HandleFunc("/teams", handlers.Teams)

	// TODO: ATM current-gameweek, teams and players endpoints all call FPL API's bootstrap-static endpoint
	// So will need to change and introduce some kind of a cache layer here for that specific call
	// ALL FPL players
	http.HandleFunc("/players", handlers.Players)

	// FPL manager and squads
	http.HandleFunc("/manager/{id}", handlers.Manager)
	http.HandleFunc("/manager/{manager_id}/squad/{gameweek_id}", handlers.Squad)
	// Required query params: gameweek, count, start
	http.HandleFunc("/manager/{id}/fixture-analysis", handlers.ManagerFixtureAnalysis)

	// Premier league player information endpoints
	http.HandleFunc("/player/{id}", handlers.Player) // General player statistics and information
	http.HandleFunc("/player/{id}/summary", handlers.PlayerSummary) // Players past and upcoming fixtures

	log.Fatal(http.ListenAndServe(":8080", nil))
}
