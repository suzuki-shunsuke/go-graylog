package handler_test

import (
	"bytes"
	"net/http"
)

type plainClient struct {
	Name     string
	Password string
}

func (client *plainClient) Get(endpoint, body string) (*http.Response, error) {
	return client.request(http.MethodGet, endpoint, body)
}

func (client *plainClient) Post(endpoint, body string) (*http.Response, error) {
	return client.request(http.MethodPost, endpoint, body)
}

func (client *plainClient) Put(endpoint, body string) (*http.Response, error) {
	return client.request(http.MethodPut, endpoint, body)
}

func (client *plainClient) Delete(endpoint, body string) (*http.Response, error) {
	return client.request(http.MethodDelete, endpoint, body)
}

func (client *plainClient) request(method, endpoint, body string) (*http.Response, error) {
	req, err := http.NewRequest(method, endpoint, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(client.Name, client.Password)
	req.Header.Set("Content-Type", "application/json")
	hc := &http.Client{}
	return hc.Do(req)
}
