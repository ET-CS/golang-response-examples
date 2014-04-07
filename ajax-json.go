// Client/Server AJAX JSON Communication using golang web-server and JQuery
// Visit: http://127.0.0.1:8080
package main

import (
    "encoding/json"
    "html/template"
    "net/http"
    "path"
)

type Data struct {
    Quote string
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fp := path.Join("templates", "ajax-json.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    //parse request to struct
    var d Data
    err := json.NewDecoder(r.Body).Decode(&d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    // create json response from struct
    a, err := json.Marshal(d)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    w.Write(a)
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/ajax", ajaxHandler)
    http.ListenAndServe(":8080", nil)
}

