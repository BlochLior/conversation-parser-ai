package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/submit", handlers.AnalyzeHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	log.Println("🔧 Go server running at :8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
