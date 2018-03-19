package testutil

import (
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/suzuki-shunsuke/go-graylog/client"
	"github.com/suzuki-shunsuke/go-graylog/mockserver"
)

// GetServerAndClient returns server and client.
// If you want to use mock server, pass "mock" as endpoint.
// If you want to use real server, pass "real" as endpoint.
// If endpoint is "" and GRAYLOG_WEB_ENDPOINT_URI is set, returns real server.
func GetServerAndClient() (*mockserver.Server, *client.Client, error) {
	var (
		server *mockserver.Server
		err    error
	)
	authName := os.Getenv("GRAYLOG_AUTH_NAME")
	authPass := os.Getenv("GRAYLOG_AUTH_PASSWORD")
	if authName == "" {
		authName = "admin"
	}
	if authPass == "" {
		authPass = "admin"
	}
	endpoint := os.Getenv("GRAYLOG_WEB_ENDPOINT_URI")
	if endpoint == "" {
		server, err = mockserver.NewServer("", nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "Failed to Get Mock Server")
		}
		endpoint = server.GetEndpoint()
	}
	client, err := client.NewClient(endpoint, authName, authPass)
	if err != nil {
		server.Close()
		return nil, nil, errors.Wrap(err, "Failed to NewClient")
	}
	if server != nil {
		server.Start()
	}
	return server, client, nil
}

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func RandStringBytesMaskImprSrc(n int) string {
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
