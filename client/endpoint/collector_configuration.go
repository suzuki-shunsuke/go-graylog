package endpoint

import (
	"fmt"
	"net/url"
)

// CollectorConfigurations returns an Collector Configuration API's endpoint url.
func (ep *Endpoints) CollectorConfigurations() string {
	// /plugins/org.graylog.plugins.collector/configurations
	return ep.collectorConfigurations.String()
}

// CollectorConfiguration returns an Collector Configuration API's endpoint url.
func (ep *Endpoints) CollectorConfiguration(id string) (*url.URL, error) {
	// /plugins/org.graylog.plugins.collector/configurations
	return urlJoin(ep.collectorConfigurations, id)
}

// CollectorConfigurationInputs returns an Collector Configuration Input API's endpoint url.
func (ep *Endpoints) CollectorConfigurationInputs(id string) (*url.URL, error) {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/inputs
	return urlJoin(ep.collectorConfigurations, fmt.Sprintf("%s/inputs", id))
}

// CollectorConfigurationInput returns an Collector Configuration Input API's endpoint url.
func (ep *Endpoints) CollectorConfigurationInput(id, inputID string) (*url.URL, error) {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId}
	return urlJoin(ep.collectorConfigurations, fmt.Sprintf("%s/inputs/%s", id, inputID))
}

// CollectorConfigurationOutputs returns an Collector Configuration Output API's endpoint url.
func (ep *Endpoints) CollectorConfigurationOutputs(id string) (*url.URL, error) {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/outputs
	return urlJoin(ep.collectorConfigurations, fmt.Sprintf("%s/outputs", id))
}

// CollectorConfigurationOutput returns an Collector Configuration Output API's endpoint url.
func (ep *Endpoints) CollectorConfigurationOutput(id, outputID string) (*url.URL, error) {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId}
	return urlJoin(ep.collectorConfigurations, fmt.Sprintf("%s/outputs/%s", id, outputID))
}

// CollectorConfigurationSnippets returns an Collector Configuration Snippet API's endpoint url.
func (ep *Endpoints) CollectorConfigurationSnippets(id string) (*url.URL, error) {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/snippets
	return urlJoin(ep.collectorConfigurations, fmt.Sprintf("%s/snippets", id))
}

// CollectorConfigurationSnippet returns an Collector Configuration Snippet API's endpoint url.
func (ep *Endpoints) CollectorConfigurationSnippet(id, snippetID string) (*url.URL, error) {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippetId}
	return urlJoin(ep.collectorConfigurations, fmt.Sprintf("%s/snippets/%s", id, snippetID))
}
