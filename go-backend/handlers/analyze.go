package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/BlochLior/conversation-parser-ai/go-backend/internal"
	"github.com/BlochLior/conversation-parser-ai/go-backend/utils"
)

var pythonAIURL = "http://localhost:8001/analyze"

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "method not allowed", nil)
		return
	}

	var req internal.AnalyzeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "invalid JSON body", err)
		return
	}

	if err = req.Validate(); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "validation failed", err)
		return
	}

	out, err := json.Marshal(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "failed to encode request", err)
		return
	}

	resp, err := http.Post(pythonAIURL, "application/json", bytes.NewReader(out))
	if err != nil {
		utils.RespondWithError(w, http.StatusBadGateway, "failed to reach AI service", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "failed to read AI service", err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("AI service returned %d: %s", resp.StatusCode, string(body))
		utils.RespondWithError(w, http.StatusBadGateway, "AI service error", nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, json.RawMessage(body))
}
