package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/kaktusas2598/FPLBuddy/backend/internal/fpl"
	"github.com/kaktusas2598/FPLBuddy/backend/internal/handlers"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	http.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		// We define a go string and then use Unmarshal to store it in json object
		graphQLQuery := `{"query":"{ teams { items { id name short_name code } } }"}`

		// Creating POST request with graphql query as a body
		req, err := http.NewRequestWithContext(
			r.Context(),
			"POST",
			"https://fpl-api-6h0d.onrender.com/graphql",
			bytes.NewBufferString(graphQLQuery),
		)
		if err != nil {
			fmt.Println("Error creating GraphQL request: ", err)
			return
		}

		// Actually sending off the request
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making GraphQL request: ", err)
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
		}

		//fmt.Fprintf(w, "%s", rawResponse)

		var teams []fpl.Team
		// Creating a nested anonymous struct to allow unmarshalling raw json response into typed structure
		var teamsResult struct {
			Data struct {
				Teams struct {
					Items []fpl.Team `json:"items"`
				} `json:"teams"`
			} `json:"data"`
		}
		err = json.Unmarshal(rawResponse, &teamsResult)
		if err != nil {
			fmt.Println("JSON Unmarshal error: ", err)
			return
		}

		//fmt.Printf("%+v", teamsResult)

		teams = teamsResult.Data.Teams.Items

		// Set response type and encode it as json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(teams)
	})

	http.HandleFunc("/current-gameweek", func(w http.ResponseWriter, r *http.Request) {
		graphQLQuery := `{"query":"{ events(where: { is_current: { eq: true } }) { items { id name deadline_time most_selected { web_name } top_element { web_name } } } }"}`

		// Creating POST request with graphql query as a body
		req, err := http.NewRequestWithContext(
			r.Context(),
			"POST",
			"https://fpl-api-6h0d.onrender.com/graphql",
			bytes.NewBufferString(graphQLQuery),
		)
		if err != nil {
			fmt.Println("Error creating GraphQL request: ", err)
			return
		}

		// Actually sending off the request
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making GraphQL request: ", err)
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
		}

		var currentGameweek fpl.GameWeek
		// Creating a nested anonymous struct to allow unmarshalling raw json response into typed structure
		var eventResult struct {
			Data struct {
				Events struct {
					Items []fpl.GameWeek `json:"items"`
				} `json:"events"`
			} `json:"data"`
		}
		err = json.Unmarshal(rawResponse, &eventResult)
		if err != nil {
			fmt.Println("JSON Unmarshal error: ", err)
			return
		}

		currentGameweek = eventResult.Data.Events.Items[0]

		// Set response type and encode it as json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(currentGameweek)
	})

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
