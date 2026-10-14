package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })

	http.HandleFunc("/api/players", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Haaland????")
    })



	http.ListenAndServe(":8080", nil)
}
