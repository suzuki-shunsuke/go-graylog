package client

import (
	"net/http"
)

// ErrorInfo represents Graylog API's error information.
type ErrorInfo struct {
	Type         string         `json:"type"`
	Message      string         `json:"message"`
	ResponseBody []byte         `json:"response_body"`
	Request      *http.Request  `json:"request"`
	Response     *http.Response `json:"response"`
}
