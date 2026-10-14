package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "bytes"
    "context"
)

// FPL Team struct
type Team struct {
	ID           int    `json:"id"`
	Name         string `json:"id"`
	ShortName    string `json:"id"`
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
		json.Unmarshal([]byte(graphQLQuery), &graphQlJsonQuery)
		// TODO: call unofficial FPL api by constructing GraphQL query and returning json here
		/*
		curl -X POST \
		  https://fpl-api-6h0d.onrender.com/graphql \
		  -H 'Content-Type: application/json' \
		  -d '{"query":"{ teams { items { id name short_name } } }"}'
		*/
        fmt.Fprintln(w, "TODO")
    })


	http.ListenAndServe(":8080", nil)
}
