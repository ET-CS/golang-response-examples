// Minimal Client/Server AJAX Communication using golang web-server and JQuery
// Visit: http://127.0.0.1:8080
package main

import (
    "html/template"
    "net/http"
    "path"
    "io/ioutil"
)

type Data struct {
    Name string
}

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
    data := Data{"World!"}

    fp := path.Join("templates", "ajax.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// AJAX Request Handler
func ajaxHandler(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadAll(r.Body);
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write(data)
    w.Write([]byte("consectetur adipisicing elit."))
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/ajax", ajaxHandler)
    http.ListenAndServe(":8080", nil)
}

