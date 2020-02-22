package main

import (
	"html/template"
	"net/http"
)

// Template context
type Context struct {
	Name string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := Context{Name: "Nicholas"}
	t, _ := template.ParseFiles("home.html")
	t.Execute(w, ctx)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8081", nil)
}
