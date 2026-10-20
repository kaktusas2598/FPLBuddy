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

	// GraphQL handlers - experimental usage
	http.HandleFunc("/teams", handlers.Teams)

	// FPL Gameweek Information
	http.HandleFunc("/current-gameweek", handlers.CurrentGameWeek)

	// FPL manager and squads
	http.HandleFunc("/manager/{id}", handlers.Manager)
	http.HandleFunc("/manager/{manager_id}/squad/{gameweek_id}", handlers.Squad)

	// Premier league players
	http.HandleFunc("/player/{id}", handlers.Player)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
