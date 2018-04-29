package client

import (
	"net/http"
)

// ErrorInfo represents Graylog API's error information.
// Basically Client methods (ex. CreateRole) returns this, but note that Response is closed.
type ErrorInfo struct {
	Type     string         `json:"type"`
	Message  string         `json:"message"`
	Request  *http.Request  `json:"request"`
	Response *http.Response `json:"response"`
}
