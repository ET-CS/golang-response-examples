// Miminal golang web-server XML response
// Visit: http://127.0.0.1:8080
package main

import (
    "net/http"
    "encoding/xml"
)

type Response struct {
    Name    string
    Hobbies []string
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    profile := Response{"ET", []string{"music", "programming"}}

    js, err := xml.MarshalIndent(profile, "", "  ")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/xml")
    w.Write(js)
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

