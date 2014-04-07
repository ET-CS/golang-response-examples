// HTML form minimal example using golang web-server
// Visit: http://127.0.0.1:8080
package main

import (
    "html/template"
    "net/http"
    "path"
)

type Data struct {
    Name string
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    user := r.FormValue("user")

    data := Data{user}

    fp := path.Join("templates", "form.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Pass template to http.ResponseWriter.
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

