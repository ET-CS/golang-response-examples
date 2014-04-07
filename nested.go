// Nested templates using golang web-server
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
    data := Data{"World!"}

    lp := path.Join("templates", "layout.html")
    fp := path.Join("templates", "content.html")

    tmpl, err := template.ParseFiles(lp, fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Pass template to http.ResponseWriter.
    if err := tmpl.Execute(w, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    // To buffer from template to string instead use:
    //buf := new(bytes.Buffer)
    //if err := tmpl.Execute(buf, data); err != nil {
    //    http.Error(w, err.Error(), http.StatusInternalServerError)
    //}
    //templateString := buf.String()
}

func main() {
    http.HandleFunc("/", defaultHandler)
    http.ListenAndServe(":8080", nil)
}

