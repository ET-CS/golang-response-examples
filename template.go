// Miminal golang web-server implementation
// Visit: http://127.0.0.1:8080/Your-name
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

    fp := path.Join("templates", "index.html")
    tmpl, err := template.ParseFiles(fp)
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

