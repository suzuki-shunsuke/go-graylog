package graylog

// APIError represents a Graylog API's error response body.
type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func NewAPIError(msg string) *APIError {
	return &APIError{Type: "ApiError", Message: msg}
}
