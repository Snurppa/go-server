package main

import (
	"html/template"
	"io/ioutil"
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
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	InfoLogger.Printf("Body of the request was: %s", bytes)
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/debug", debug).
		Methods("POST")
	//r.HandleFunc("/api", api)
	return r
}
