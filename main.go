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
	count := cfg.fileserverHits.Load()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d", count)))
}

func (cfg *apiConfig)resetHandler(w http.ResponseWriter,r *http.Request){
	cfg.fileserverHits.Store(0)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Counter reset"))
}

func main() {
	const port = "8080"
	const filepathroot = "."

	mux := http.NewServeMux()

	// health endpoint
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	apiConfig := &apiConfig{}
	
	// static file server
	fs := http.FileServer(http.Dir(filepathroot))
	mux.Handle("/app/", apiConfig.middlewareMetricsInc(http.StripPrefix("/app/", fs)))
	mux.HandleFunc("/metrics",apiConfig.metricsHandler)
	mux.HandleFunc("/reset",apiConfig.resetHandler)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving from %s on port: %s\n", filepathroot, port)
	log.Fatal(srv.ListenAndServe())
}
