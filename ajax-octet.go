// Minimal Client/Server AJAX Communication using golang web-server and JQuery
// Visit: http://127.0.0.1:8080
package main

import (
    "encoding/json"
    "html/template"
    "net/http"
    "path"
)

type Data struct {
    Name string
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {


    fp := path.Join("templates", "ajax-octet.html")
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
    data := Data{"World!"}
    w.Header().Set("Content-type", "application/json")
    err := json.NewEncoder(w).Encode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    return
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/ajax", ajaxHandler)
    http.ListenAndServe(":8080", nil)
}