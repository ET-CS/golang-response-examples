// Redirect response using golang web-server
// Visit: http://127.0.0.1:8080/redirect
package main

import (
    "net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("OK"))
}

// Redirect Request Handler
func redirectHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/redirect", redirectHandler)
    http.ListenAndServe(":8080", nil)
}

