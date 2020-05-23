package main

import (
	"html/template"
	"log"
	"net/http"
)

// Page Response page
type Page struct {
	Title   string
	Heading string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))
var defaultPageData = Page{Title: "jonih.fi", Heading: "Hello world from Go!"}

func handler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}
	err := templates.ExecuteTemplate(writer, "index.html", defaultPageData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	//http.Handle()
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
