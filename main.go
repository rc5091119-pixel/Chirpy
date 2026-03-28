package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}
func main() {
	const port = "8080"
	const filepathroot = "."

	mux := http.NewServeMux()

	// health endpoint
	apiConfig := &apiConfig{
		fileserverHits: atomic.Int32{},
	}

	// static file server
	fs := http.FileServer(http.Dir(filepathroot))
	mux.Handle("/app/", apiConfig.middlewareMetricsInc(http.StripPrefix("/app/", fs)))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp",apiConfig.handlerValidate_chrip)
	mux.HandleFunc("GET /admin/metrics", apiConfig.handlerMetrics)
	mux.HandleFunc("POST /admin/reset", apiConfig.handlerReset)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving from %s on port: %s\n", filepathroot, port)
	log.Fatal(srv.ListenAndServe())
}
