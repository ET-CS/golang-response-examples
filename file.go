// Serving file using golang web-server
// Visit: http://127.0.0.1:8080
package main

import (
    "net/http"
    "path"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fp := path.Join("images", "Golang.png")
    http.ServeFile(w, r, fp)
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

