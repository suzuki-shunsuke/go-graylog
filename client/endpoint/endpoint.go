package endpoint

import (
	"errors"
	"strings"
)

// Endpoints represents each API's endpoint URLs.
type Endpoints struct {
	alarmCallbacks           string
	alerts                   string
	alertConditions          string
	collectorConfigurations  string
	dashboards               string
	enabledStreams           string
	indexSets                string
	indexSetStats            string
	inputs                   string
	outputs                  string
	availableOutputs         string
	pipelines                string
	pipelineConnections      string
	pipelineRules            string
	roles                    string
	streams                  string
	users                    string
	grokPatterns             string
	grokPatternsTest         string
	ldapSetting              string
	ldapGroups               string
	ldapGroupRoleMapping     string
	connectStreamsToPipeline string
	connectPipelinesToStream string
	apiVersion               string
}

// NewEndpoints returns a new Endpoints.
func NewEndpoints(endpoint string) (*Endpoints, error) {
	return newEndpoints(endpoint, "")
}

// NewEndpointsV3 returns a new Endpoints for Graylog API v3.
func NewEndpointsV3(endpoint string) (*Endpoints, error) {
	return newEndpoints(endpoint, "v3")
}

func newEndpoints(endpoint, version string) (*Endpoints, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint is required")
	}
	endpoint = strings.TrimRight(endpoint, "/")

	var pipelines, pipelineRules, pipelineConns, connectPipelinesToStream, connectStreamsToPipeline string
	if version == "v3" {
		// https://docs.graylog.org/en/latest/pages/upgrade/graylog-3.0.html#plugins-merged-into-the-graylog-server
		pipelines = endpoint + "/system/pipelines/pipeline"
		pipelineRules = endpoint + "/system/pipelines/rule"
		pipelineConns = endpoint + "/system/pipelines/connections"
		connectStreamsToPipeline = endpoint + "/system/pipelines/connections/to_pipeline"
		connectPipelinesToStream = endpoint + "/system/pipelines/connections/to_stream"
	} else {
		pipelines = endpoint + "/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/pipeline"
		pipelineRules = endpoint + "/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/rule"
		pipelineConns = endpoint + "/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections"
		connectStreamsToPipeline = endpoint + "/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections/to_pipeline"
		connectPipelinesToStream = endpoint + "/plugins/org.graylog.plugins.pipelineprocessor/system/pipelines/connections/to_stream"
	}
	return &Endpoints{
		alarmCallbacks:          endpoint + "/alerts/callbacks",
		alerts:                  endpoint + "/streams/alerts",
		alertConditions:         endpoint + "/alerts/conditions",
		collectorConfigurations: endpoint + "/plugins/org.graylog.plugins.collector/configurations",
		dashboards:              endpoint + "/dashboards",
		enabledStreams:          endpoint + "/streams/enabled",
		indexSets:               endpoint + "/system/indices/index_sets",
		indexSetStats:           endpoint + "/system/indices/index_sets/stats",
		inputs:                  endpoint + "/system/inputs",
		ldapGroups:              endpoint + "/system/ldap/groups",
		ldapGroupRoleMapping:    endpoint + "/system/ldap/settings/groups",
		ldapSetting:             endpoint + "/system/ldap/settings",

		outputs:          endpoint + "/system/outputs",
		availableOutputs: endpoint + "/system/outputs/available",

		pipelines:                pipelines,
		pipelineConnections:      pipelineConns,
		connectStreamsToPipeline: connectStreamsToPipeline,
		connectPipelinesToStream: connectPipelinesToStream,
		pipelineRules:            pipelineRules,
		roles:                    endpoint + "/roles",
		streams:                  endpoint + "/streams",
		users:                    endpoint + "/users",
		grokPatterns:             endpoint + "/system/grok",
		grokPatternsTest:         endpoint + "/system/grok/test",
		apiVersion:               version,
	}, nil
}
