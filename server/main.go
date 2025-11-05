package main

import (
    "log"
    "net/http"
)

func main() {
    // Serve static files from the React build
    fs := http.FileServer(http.Dir("../frontend/build"))
    http.Handle("/", fs)

    log.Println("Server starting on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}