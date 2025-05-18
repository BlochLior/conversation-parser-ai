package client

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BlochLior/conversation-parser-ai/go-backend/internal"
)

func TestAnalyzeConversation_Success(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := internal.AnalyzeResponse{
			Issues:      []string{"issue1"},
			Suggestions: []string{"suggestion1"},
		}
		_ = json.NewEncoder(w).Encode(resp)
	}))
	defer mock.Close()

	cli := New(mock.URL)
	out, err := cli.AnalyzeConversation(internal.AnalyzeRequest{Conversation: "Hello"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Issues) != 1 || len(out.Suggestions) != 1 {
		t.Errorf("unexpected output: %+v", out)
	}
}

func TestAnalyzeConversation_InvalidJSON(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("{invalid-json}")); err != nil {
			log.Printf("error writing response: %v", err)
		}
	}))
	defer mock.Close()

	cli := New(mock.URL)
	_, err := cli.AnalyzeConversation(internal.AnalyzeRequest{Conversation: "Hi"})
	if err == nil {
		t.Fatal("expected decode error, got nil")
	}
}

func TestAnalyzeConversation_BadStatus(t *testing.T) {
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))
	defer mock.Close()

	cli := New(mock.URL)
	_, err := cli.AnalyzeConversation(internal.AnalyzeRequest{Conversation: "Hi"})
	if err == nil {
		t.Fatal("expected HTTP status error, got nil")
	}
}
