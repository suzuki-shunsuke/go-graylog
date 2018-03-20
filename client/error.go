package client

import (
	"net/http"
)

// ErrorInfo represents Graylog API's error information.
type ErrorInfo struct {
	Type     string         `json:"type"`
	Message  string         `json:"message"`
	Request  *http.Request  `json:"request"`
	Response *http.Response `json:"response"`
}
