package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
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
	mux.HandleFunc("GET /healthz", handlerReadiness)
	mux.HandleFunc("GET /metrics", apiConfig.metricsHandler)
	mux.HandleFunc("POST /reset", apiConfig.handlerReset)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving from %s on port: %s\n", filepathroot, port)
	log.Fatal(srv.ListenAndServe())
}
