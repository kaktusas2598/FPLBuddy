package main

import (
    "fmt"
    "net/http"
)

func main() {
	http.ListenAndServe(":8080", nil)

    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })
}
