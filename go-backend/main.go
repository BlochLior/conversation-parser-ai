package main

import (
	"log"
	"net/http"

	"github.com/BlochLior/conversation-parser-ai/go-backend/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/submit", handlers.AnalyzeHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	log.Println("🔧 Go server running at :8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
