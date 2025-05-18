package main

import (
	"log"
	"net/http"
	"time"

	"github.com/BlochLior/conversation-parser-ai/go-backend/handlers"
	"github.com/BlochLior/conversation-parser-ai/go-backend/utils"
	"github.com/BlochLior/conversation-parser-ai/shared/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/submit", handlers.AnalyzeHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      cors.WithCORS(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	log.Println("🔧 Go server running at :8000")

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
