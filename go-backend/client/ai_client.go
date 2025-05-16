package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BlochLior/conversation-parser-ai/go-backend/internal"
)

// AIClient represents the service that communicates with the Python AI API
type AIClient struct {
	BaseURL string
	Client  *http.Client
}

func New(baseURL string) *AIClient {
	return &AIClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func (c *AIClient) AnalyzeConversation(req internal.AnalyzeRequest) (*internal.AnalyzeResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("marshal error: %w", err)
	}

	res, err := c.Client.Post(c.BaseURL+"/analyze", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("AI service returned status: %s", res.Status)
	}

	var out internal.AnalyzeResponse
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}

	return &out, nil
}
