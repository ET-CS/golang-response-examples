// Miminal golang web-server JSON response
// Visit: http://127.0.0.1:8080
package main

import (
    "net/http"
    "encoding/json"
)

// Maybe you want to struct your JSON in a specification? 
// http://labs.omniti.com/labs/jsend
type Response struct {
    Name    string
    Hobbies []string
    // remember: Go enforce you to use uppercase first letter in each field for json.Marshal.
    // If you want to use lower case json response, you can tag the fields like:
    // Hobbies []string `json:"hobbies"`
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    profile := Response{"ET", []string{"music", "programming"}}

    js, err := json.Marshal(profile)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

