package main

import (
	"encoding/binary"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"raspi/server/db"

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
	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		ErrorLogger.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	value, error := binary.Uvarint(b)
	if error <= 0 {
		ErrorLogger.Printf("Failed to read uint from request bytes: %d", error)
	} else {
		db.WriteMoisture(os.Getenv("INFLUX_ORG"), os.Getenv("INFLUX_BUCKET"), "peikko", int(value))
	}
	InfoLogger.Printf("Body of the request was: %s", b)
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/debug", debug).
		Methods("POST")
	//r.HandleFunc("/api", api)
	return r
}
