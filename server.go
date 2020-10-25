package main

import (
	"log"
	"net/http"
)

func main() {
	InfoLogger.Println("Bootstrapping")
	http.Handle("/", router())
	InfoLogger.Println("Router ready, starting server")
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}
