package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/BlochLior/conversation-parser-ai/go-backend/client"
	"github.com/BlochLior/conversation-parser-ai/go-backend/internal"
	"github.com/BlochLior/conversation-parser-ai/go-backend/utils"
)

var aiService = client.New("http://localhost:8001")

// var pythonAIURL = "http://localhost:8001/analyze"

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

	resp, err := aiService.AnalyzeConversation(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadGateway, "AI service error", err)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, resp)
}
