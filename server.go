package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	host := ""
	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	}
	port := "8888"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	InfoLogger.Printf("Bootstrapping system on %s\n", runtime.GOARCH)
	http.Handle("/", router())
	InfoLogger.Printf("Router ready, starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
