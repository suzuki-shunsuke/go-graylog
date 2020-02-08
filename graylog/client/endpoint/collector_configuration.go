package endpoint

// CollectorConfigurations returns a Collector Configuration API's endpoint url.
func (ep *Endpoints) CollectorConfigurations() string {
	// /plugins/org.graylog.plugins.collector/configurations
	return ep.collectorConfigurations
}

// CollectorConfiguration returns a Collector Configuration API's endpoint url.
func (ep *Endpoints) CollectorConfiguration(id string) string {
	// /plugins/org.graylog.plugins.collector/configurations
	return ep.collectorConfigurations + "/" + id
}

// CollectorConfigurationName returns a Collector Configuration API's endpoint url.
func (ep *Endpoints) CollectorConfigurationName(id string) string {
	// /plugins/org.graylog.plugins.collector/configurations/:id/name
	return ep.collectorConfigurations + "/" + id + "/name"
}

// CollectorConfigurationInputs returns a Collector Configuration Input API's endpoint url.
func (ep *Endpoints) CollectorConfigurationInputs(id string) string {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/inputs
	return ep.collectorConfigurations + "/" + id + "/inputs"
}

// CollectorConfigurationInput returns a Collector Configuration Input API's endpoint url.
func (ep *Endpoints) CollectorConfigurationInput(id, inputID string) string {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/inputs/{inputId}
	return ep.collectorConfigurations + "/" + id + "/inputs/" + inputID
}

// CollectorConfigurationOutputs returns a Collector Configuration Output API's endpoint url.
func (ep *Endpoints) CollectorConfigurationOutputs(id string) string {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/outputs
	return ep.collectorConfigurations + "/" + id + "/outputs"
}

// CollectorConfigurationOutput returns a Collector Configuration Output API's endpoint url.
func (ep *Endpoints) CollectorConfigurationOutput(id, outputID string) string {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/outputs/{outputId}
	return ep.collectorConfigurations + "/" + id + "/outputs/" + outputID
}

// CollectorConfigurationSnippets returns a Collector Configuration Snippet API's endpoint url.
func (ep *Endpoints) CollectorConfigurationSnippets(id string) string {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/snippets
	return ep.collectorConfigurations + "/" + id + "/snippets"
}

// CollectorConfigurationSnippet returns a Collector Configuration Snippet API's endpoint url.
func (ep *Endpoints) CollectorConfigurationSnippet(id, snippetID string) string {
	// /plugins/org.graylog.plugins.collector/configurations/{id}/snippets/{snippetId}
	return ep.collectorConfigurations + "/" + id + "/snippets/" + snippetID
}
