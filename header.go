// Set header on golang web-server
// Try: curl -i localhost:8080
package main

import (
    "net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Server", "A Go Web Server")
    w.WriteHeader(200)
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

