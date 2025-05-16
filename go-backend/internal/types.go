package internal

import "errors"

// AnalyzeRequest matches the expected JSON body sent to the Python AI service
type AnalyzeRequest struct {
	Conversation string `json:"conversation"`
}

func (a AnalyzeRequest) Validate() error {
	if a.Conversation == "" {
		return errors.New("conversation is required")
	}
	return nil
}

// AnalyzeResponse matches the JSON response returned by the Python AI service
type AnalyzeResponse struct {
	Issues      []string `json:"issues"`
	Suggestions []string `json:"suggestions"`
}
