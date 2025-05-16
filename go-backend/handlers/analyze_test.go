package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BlochLior/conversation-parser-ai/go-backend/client"
)

func TestAnalyzeHandler_Success(t *testing.T) {
	mockAI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"issues": ["tone mismatch"], "suggestions": ["clarify intent"]}`))
	}))
	defer mockAI.Close()

	aiService = client.New(mockAI.URL)

	payload := `{"conversation": "Speaker A: Hello\nSpeaker B: What?"}`
	r := httptest.NewRequest("POST", "/submit", bytes.NewBuffer([]byte(payload)))
	w := httptest.NewRecorder()

	AnalyzeHandler(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", w.Code)
	}

	expected := `{"issues":["tone mismatch"],"suggestions":["clarify intent"]}`
	if w.Body.String() != expected {
		t.Errorf("unexpected body: %s", w.Body.String())
	}
}

func TestAnalyzeHandler_EmptyConversation(t *testing.T) {
	payload := `{"conversation": ""}`
	r := httptest.NewRequest("POST", "/submit", bytes.NewBuffer([]byte(payload)))
	w := httptest.NewRecorder()

	AnalyzeHandler(w, r)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 Bad Request, got %d", w.Code)
	}
}
