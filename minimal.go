// Miminal golang web-server implementation
// Visit: http://127.0.0.1:8080/Your-name
package main

import (
    "net/http"
    "fmt"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    // using fmt.fprintf()
    fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
    // using w.Write()
    w.Write([]byte("OK"))
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

