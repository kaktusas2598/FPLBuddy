package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/handlers"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	http.HandleFunc("/current-gameweek", handlers.CurrentGameWeek)
	http.HandleFunc("/teams", handlers.Teams)


	http.HandleFunc("/manager/{id}", func(w http.ResponseWriter, r *http.Request) {
		managerID := r.PathValue("id")

		url := fmt.Sprintf(
			"https://fantasy.premierleague.com/api/entry/%s/",
			managerID,
		)

		// Unofficial FPL GraphQL does not support fetching the squad data so here we are using the official API
		req, err := http.NewRequestWithContext(r.Context(), "GET", url, nil)
		if err != nil {
			fmt.Println("Error creating FPL API request: ", err)
			return
		}

		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making FPL API request: ", err)
			return
		}

		defer response.Body.Close()

		if response.StatusCode != 200 {
			fmt.Println("response status: ", response.StatusCode)
			return
		}

		rawResponse, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(rawResponse)
	})

	http.HandleFunc("/player/{id}", handlers.Player)

	http.ListenAndServe(":8080", nil)
}
