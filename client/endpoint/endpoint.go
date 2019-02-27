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
	alarmCallbacks          *url.URL
	alerts                  *url.URL
	alertConditions         *url.URL
	collectorConfigurations *url.URL
	dashboards              *url.URL
	enabledStreams          *url.URL
	indexSets               *url.URL
	indexSetStats           *url.URL
	inputs                  *url.URL
	pipelineRules           *url.URL
	roles                   *url.URL
	streams                 *url.URL
	users                   *url.URL
	ldapSetting             string
	ldapGroups              string
	ldapGroupRoleMapping    string
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
	alarmCallbacks, err := urlJoin(ep, "alerts/callbacks")
	if err != nil {
		return nil, err
	}
	alertConditions, err := urlJoin(ep, "alerts/conditions")
	if err != nil {
		return nil, err
	}
	collectorConfigurations, err := urlJoin(ep, "plugins/org.graylog.plugins.collector/configurations")
	if err != nil {
		return nil, err
	}
	dashboards, err := urlJoin(ep, "dashboards")
	if err != nil {
		return nil, err
	}
	roles, err := urlJoin(ep, "roles")
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
	indexSets, err := urlJoin(ep, "system/indices/index_sets")
	if err != nil {
		return nil, err
	}
	indexSetStats, err := urlJoin(indexSets, "stats")
	if err != nil {
		return nil, err
	}
	inputs, err := urlJoin(ep, "system/inputs")
	if err != nil {
		return nil, err
	}
	ldapSetting, err := urlJoin(ep, "system/ldap/settings")
	if err != nil {
		return nil, err
	}
	ldapGroups, err := urlJoin(ep, "system/ldap/groups")
	if err != nil {
		return nil, err
	}
	ldapGroupRoleMapping, err := urlJoin(ep, "system/ldap/settings/groups")
	if err != nil {
		return nil, err
	}
	pipelineRules, err := urlJoin(ep, "plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/rule")
	if err != nil {
		return nil, err
	}
	users, err := urlJoin(ep, "users")
	if err != nil {
		return nil, err
	}
	return &Endpoints{
		alarmCallbacks:          alarmCallbacks,
		alerts:                  alerts,
		alertConditions:         alertConditions,
		collectorConfigurations: collectorConfigurations,
		dashboards:              dashboards,
		enabledStreams:          enabledStreams,
		indexSets:               indexSets,
		indexSetStats:           indexSetStats,
		inputs:                  inputs,
		ldapGroups:              ldapGroups.String(),
		ldapGroupRoleMapping:    ldapGroupRoleMapping.String(),
		ldapSetting:             ldapSetting.String(),
		pipelineRules:           pipelineRules,
		roles:                   roles,
		streams:                 streams,
		users:                   users,
	}, nil
}
