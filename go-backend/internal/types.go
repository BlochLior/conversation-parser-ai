package internal

// AnalyzeRequest matches the expected JSON body sent to the Python AI service
type AnalyzeRequest struct {
	Conversation string `json:"conversation"`
}

// AnalyzeResponse matches the JSON response returned by the Python AI service
type AnalyzeResponse struct {
	Issues      []string `json:"issues"`
	Suggestions []string `json:"suggestions"`
}
