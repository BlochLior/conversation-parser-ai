package main

import (
	"log"
	"net/http"
	"time"

	"github.com/BlochLior/conversation-parser-ai/shared/cors"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("ok")); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.Handle("/", http.FileServer(http.Dir("./static")))

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      cors.WithCORS(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("🌐 Frontend available at http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("❌ Failed to start frontend: %v", err)
	}
}
