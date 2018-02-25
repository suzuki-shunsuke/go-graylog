package graylog

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func getServerAndClient() (*MockServer, *Client, error) {
	server, err := NewMockServer("")
	if err != nil {
		return nil, nil, errors.Wrap(err, "Failed to Get Mock Server")
	}
	client, err := NewClient(server.Endpoint, "admin", "password")
	if err != nil {
		server.Server.Close()
		return nil, nil, errors.Wrap(err, "Failed to NewClient")
	}
	server.Start()
	return server, client, nil
}

func callRequest(
	req *http.Request, client *Client, ctx context.Context,
) (*http.Response, error) {
	req.SetBasicAuth(client.GetName(), client.GetPassword())
	req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	hc := &http.Client{}
	return hc.Do(req)
}

func addToStringArray(arr []string, val string) []string {
	for _, v := range arr {
		if v == val {
			return arr
		}
	}
	return append(arr, val)
}

func removeFromStringArray(arr []string, val string) []string {
	ret := []string{}
	for _, v := range arr {
		if v != val {
			ret = append(ret, v)
		}
	}
	return ret
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func randStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
