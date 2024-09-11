package model

// ErrorResponse represents an error response that can be returned in JSON format.
type ErrorResponse struct {
	// Error is the human-readable error message.
	Error string `json:"error,omitempty"`
}
