package handler

// APIError represents a Graylog API's error response body.
type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// NewAPIError returns a new APIError.
func NewAPIError(msg string) *APIError {
	return &APIError{Type: "ApiError", Message: msg}
}
