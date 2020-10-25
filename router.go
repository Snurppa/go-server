package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// Page Response page
type Page struct {
	Title   string
	Heading string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))
var defaultPageData = Page{Title: "RaspiHome", Heading: "Hello world from Go!"}

func root(writer http.ResponseWriter, request *http.Request) {
	err := templates.ExecuteTemplate(writer, "index.html", defaultPageData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func debug(writer http.ResponseWriter, request *http.Request) {
	InfoLogger.Println(request)
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/debug", debug).
		Methods("POST")
	//r.HandleFunc("/api", api)
	return r
}
