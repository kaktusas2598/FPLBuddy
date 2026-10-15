package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "bytes"
    "io"
)

// FPL Team struct
type Team struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ShortName    string `json:"short_name"`
}

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })

	http.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Haaland????")
    })

	http.HandleFunc("/teams", func(w http.ResponseWriter, r *http.Request) {
		// We define a go string and then use Unmarshal to store it in json object
		graphQLQuery := `{"query":"{ teams { items { id name short_name } } }"}`

		// Creating POST request with graphql query as a body
		req, err := http.NewRequestWithContext(
			r.Context(),
			"POST",
			"https://fpl-api-6h0d.onrender.com/graphql",
			bytes.NewBufferString(graphQLQuery),
		)
		req.Header.Set("Content-Type", "application/json")

		if err != nil {
			fmt.Println("Error creating GraphQL request: ", err)
			return
		}

		// Actually sending off the request
		client := &http.Client{}
		response, err := client.Do(req)

		if err != nil {
			fmt.Println("Error making GraphQL request: ", err)
			return
		}

		defer response.Body.Close()


		fmt.Println("Request response status: ", response.StatusCode)

		rawResponse, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		//fmt.Fprintf(w, "%s", rawResponse)


		var teams []Team
		// Creating a nested anonymous struct to allow unmarshalling raw json response into typed structure
		var teamsResult struct {
			Data struct {
				Teams struct {
					Items []Team `json:"items"`
				} `json:"teams"`
			} `json:"data"`
		}
		err = json.Unmarshal(rawResponse, &teamsResult)

		if err != nil {
			fmt.Println("JSON Unmarshal error: ", err)
		}

		//fmt.Printf("%+v", teamsResult)

		teams = teamsResult.Data.Teams.Items


		fmt.Fprintln(w, teams[5].Name)
    })


	http.ListenAndServe(":8080", nil)
}
