package graylog

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (client *Client) callReq(
	ctx context.Context, method, endpoint string,
	body []byte, isReadBody bool,
) (*ErrorInfo, error) {
	var reqBody io.Reader = nil
	if body != nil {
		reqBody = bytes.NewBuffer(body)
	}
	req, err := http.NewRequest(method, endpoint, reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to http.NewRequest")
	}
	ei := &ErrorInfo{Request: req}
	req.SetBasicAuth(client.GetName(), client.GetPassword())
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		return ei, errors.Wrap(
			err, fmt.Sprintf("Failed to call Graylog API: %s %s", method, endpoint))
	}
	ei.Response = resp

	if resp.StatusCode >= 400 {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ei, errors.Wrap(err, "Failed to read response body")
		}
		ei.ResponseBody = b
		if err := json.Unmarshal(b, ei); err != nil {
			return ei, errors.Wrap(
				err, fmt.Sprintf(
					"Failed to parse response body as ErrorInfo: %s", string(b)))
		}
		return ei, errors.New(ei.Message)
	}

	if isReadBody {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ei, errors.Wrap(err, "Failed to read response body")
		}
		ei.ResponseBody = b
	}
	return ei, nil
}
