package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {

	data := PageData {
		Title: "TODO List",
		Todos: []Todo{
			{Item: "Install GO", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Like this video", Done: false},
		},
	}

	tmpl.Execute(w, data)

}


func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/todo", todo)
	
	log.Fatal(http.ListenAndServe("0.0.0.0:9091", mux))
}
