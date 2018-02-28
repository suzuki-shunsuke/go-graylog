package graylog

// Error represents Graylog API's error response body.
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
