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
	http.HandleFunc("/current-gameweek", handlers.CurrentGameWeek)
	http.HandleFunc("/teams", handlers.Teams)

	// Official FPL API based handlers
	http.HandleFunc("/manager/{id}", handlers.Manager)
	http.HandleFunc("/player/{id}", handlers.Player)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
