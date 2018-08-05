package endpoint

import (
	"fmt"
	"net/url"
	"path"
)

func urlJoin(ep *url.URL, arg string) (*url.URL, error) {
	return ep.Parse(path.Join(ep.Path, arg))
}

// Endpoints represents each API's endpoint URLs.
type Endpoints struct {
	roles           *url.URL
	users           *url.URL
	inputs          *url.URL
	indexSets       *url.URL
	indexSetStats   *url.URL
	streams         *url.URL
	enabledStreams  *url.URL
	alerts          *url.URL
	alertConditions *url.URL
	dashboards      *url.URL
}

// NewEndpoints returns a new Endpoints.
func NewEndpoints(endpoint string) (*Endpoints, error) {
	if endpoint == "" {
		return nil, fmt.Errorf("endpoint is required")
	}
	ep, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	roles, err := urlJoin(ep, "roles")
	if err != nil {
		return nil, err
	}
	users, err := urlJoin(ep, "users")
	if err != nil {
		return nil, err
	}
	inputs, err := urlJoin(ep, "system/inputs")
	if err != nil {
		return nil, err
	}
	indexSets, err := urlJoin(ep, "system/indices/index_sets")
	if err != nil {
		return nil, err
	}
	indexSetStats, err := urlJoin(indexSets, "stats")
	if err != nil {
		return nil, err
	}
	streams, err := urlJoin(ep, "streams")
	if err != nil {
		return nil, err
	}
	enabledStreams, err := urlJoin(streams, "enabled")
	if err != nil {
		return nil, err
	}
	alerts, err := urlJoin(ep, "streams/alerts")
	if err != nil {
		return nil, err
	}
	alertConditions, err := urlJoin(ep, "alerts/conditions")
	if err != nil {
		return nil, err
	}
	dashboards, err := urlJoin(ep, "dashboards")
	if err != nil {
		return nil, err
	}
	return &Endpoints{
		roles:           roles,
		users:           users,
		inputs:          inputs,
		indexSets:       indexSets,
		indexSetStats:   indexSetStats,
		streams:         streams,
		enabledStreams:  enabledStreams,
		alerts:          alerts,
		alertConditions: alertConditions,
		dashboards:      dashboards,
	}, nil
}
