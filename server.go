package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8888"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	InfoLogger.Println("Bootstrapping")
	http.Handle("/", router())
	InfoLogger.Printf("Router ready, starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
