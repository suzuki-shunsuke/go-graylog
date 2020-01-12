package graylog

import (
	"encoding/json"
	"net/http"
	"sync"
	"testing"

	"github.com/hashicorp/terraform/terraform"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/suzuki-shunsuke/flute/flute"
)

func genTestBody(exp map[string]interface{}, bodyString string, store *bodyStore) func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
	return func(t *testing.T, req *http.Request, svc *flute.Service, route *flute.Route) {
		body := map[string]interface{}{}
		require.Nil(t, json.NewDecoder(req.Body).Decode(&body))
		assert.Equal(t, exp, body)
		store.Set(bodyString)
	}
}

func getTestProviders() map[string]terraform.ResourceProvider {
	return map[string]terraform.ResourceProvider{
		"graylog": Provider(),
	}
}

func getTestHeader() http.Header {
	return http.Header{
		"Content-Type":   []string{"application/json"},
		"X-Requested-By": []string{"terraform-provider-graylog"},
		"Authorization":  nil,
	}
}

type bodyStore struct {
	body  string
	mutex *sync.RWMutex
}

func newBodyStore(body string) *bodyStore {
	return &bodyStore{
		body:  body,
		mutex: &sync.RWMutex{},
	}
}

func (store *bodyStore) Get() string {
	store.mutex.RLock()
	a := store.body
	store.mutex.RUnlock()
	return a
}

func (store *bodyStore) Set(body string) {
	store.mutex.Lock()
	store.body = body
	store.mutex.Unlock()
}
