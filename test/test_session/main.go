package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/", http.HandlerFunc(home))

	http.ListenAndServe(":5000", nil)
}



func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

