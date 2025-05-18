package main

import (
	"log"
	"net/http"
	"time"

	"github.com/BlochLior/conversation-parser-ai/shared/cors"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	mux := http.NewServeMux()
	mux.Handle("/", fs)

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
