package main

import {
    "fmt"
    "log"
    "net/http"
}

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "OK")
    })
}
