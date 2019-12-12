package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (client *Client) callGet(
	ctx context.Context, endpoint string, input, output interface{}) (*ErrorInfo, error) {
	return client.callAPI(ctx, http.MethodGet, endpoint, input, output)
}

func (client *Client) callPost(
	ctx context.Context, endpoint string, input, output interface{}) (*ErrorInfo, error) {
	return client.callAPI(ctx, http.MethodPost, endpoint, input, output)
}

func (client *Client) callPut(
	ctx context.Context, endpoint string, input, output interface{}) (*ErrorInfo, error) {
	return client.callAPI(ctx, http.MethodPut, endpoint, input, output)
}

func (client *Client) callDelete(
	ctx context.Context, endpoint string, input, output interface{}) (*ErrorInfo, error) {
	return client.callAPI(ctx, http.MethodDelete, endpoint, input, output)
}

func (client *Client) callAPI(
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
			return nil, errors.Wrap(err, "failed to encode request body")
		}
		req, err = http.NewRequest(method, endpoint, reqBody)
	} else {
		req, err = http.NewRequest(method, endpoint, nil)
	}
	if err != nil {
		return nil, errors.Wrapf(
			err, "failed to call http.NewRequest: %s %s", method, endpoint)
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
		return ei, errors.Wrapf(
			err, "failed to call Graylog API: %s %s", method, endpoint)
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
			return ei, errors.Wrapf(
				err, "failed to parse response body as ErrorInfo: %s %s %d "+string(b),
				method, endpoint, resp.StatusCode)
		}
		return ei, fmt.Errorf(
			"graylog API error: %s %s %d: "+string(b),
			method, endpoint, resp.StatusCode)
	}
	if output != nil {
		if err := json.NewDecoder(ei.Response.Body).Decode(output); err != nil {
			return ei, errors.Wrapf(
				err, "failed to decode graylog API response body: %s %s",
				method, endpoint)
		}
	}
	return ei, nil
}
