package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	const filepathroot = "."

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(filepathroot))
	mux.Handle("/", fs)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving from %s on port: %s\n", filepathroot, port)
	log.Fatal(srv.ListenAndServe())

}
