package client_test

import (
	"net/http"
)

func getTestHeader() http.Header {
	return http.Header{
		"Content-Type":   []string{"application/json"},
		"X-Requested-By": []string{"go-graylog"},
		"Authorization":  nil,
	}
}
