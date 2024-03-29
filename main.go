package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hello, World!</h1>")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server is listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
