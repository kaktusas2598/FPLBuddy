package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "bytes"
    "context"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })

	http.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Haaland????")
    })

	http.HandleFunc("/players", func(w http.ResponseWriter, r *http.Request) {
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
