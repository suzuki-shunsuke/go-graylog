package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (client *Client) callGet(
	ctx context.Context, endpoint string, input, output interface{}) (*ErrorInfo, error) {
	return client.CallAPI(ctx, http.MethodGet, endpoint, input, output)
}

func (client *Client) callPost(
	ctx context.Context, endpoint string, input, output interface{}) (*ErrorInfo, error) {
	return client.CallAPI(ctx, http.MethodPost, endpoint, input, output)
}

func (client *Client) callPut(
	ctx context.Context, endpoint string, input, output interface{}) (*ErrorInfo, error) {
	return client.CallAPI(ctx, http.MethodPut, endpoint, input, output)
}

func (client *Client) callDelete(
	ctx context.Context, endpoint string, input, output interface{}) (*ErrorInfo, error) {
	return client.CallAPI(ctx, http.MethodDelete, endpoint, input, output)
}

func (client *Client) CallAPI(
	ctx context.Context, method, endpoint string, input, output interface{},
) (*ErrorInfo, error) {
	// prepare request
	var (
		req *http.Request
		err error
	)
	if input != nil {
		reqBody := &bytes.Buffer{}
		if err := json.NewEncoder(reqBody).Encode(input); err != nil {
			return nil, fmt.Errorf("failed to encode request body: %w", err)
		}
		req, err = http.NewRequest(method, endpoint, reqBody)
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}
	if err != nil {
		return nil, fmt.Errorf(
			"failed to call http.NewRequest: %s %s: %w", method, endpoint, err)
	}
	ei := &ErrorInfo{Request: req}
	req.SetBasicAuth(client.Name(), client.Password())
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	// https://github.com/suzuki-shunsuke/go-graylog/issues/42
	req.Header.Set("X-Requested-By", client.xRequestedBy)
	hc := client.httpClient
	if hc == nil {
		hc = http.DefaultClient
	}
	// request
	resp, err := hc.Do(req)
	if err != nil {
		return ei, fmt.Errorf(
			"failed to call Graylog API: %s %s: %w", method, endpoint, err)
	}
	defer resp.Body.Close()
	ei.Response = resp

	if resp.StatusCode >= 400 {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ei, fmt.Errorf(
				"graylog API error: failed to read the response body: %s %s %d",
				method, endpoint, resp.StatusCode)
		}
		if err := json.Unmarshal(b, ei); err != nil {
			return ei, fmt.Errorf(
				"failed to parse response body as ErrorInfo: %s %s %d %s: %w",
				method, endpoint, resp.StatusCode, string(b), err)
		}
		return ei, fmt.Errorf(
			"graylog API error: %s %s %d: "+string(b),
			method, endpoint, resp.StatusCode)
	}
	if output != nil {
		if err := json.NewDecoder(ei.Response.Body).Decode(output); err != nil {
			return ei, fmt.Errorf(
				"failed to decode graylog API response body: %s %s: %w",
				method, endpoint, err)
		}
	}
	return ei, nil
}
