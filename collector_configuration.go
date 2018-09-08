package graylog

type (
	// CollectorConfiguration represents a Graylog's Collector Configuration.
	CollectorConfiguration struct {
		ID   string `json:"id,omitempty" v-create:"isdefault"`
		Name string `json:"name,omitempty" v-create:"required"`
		// Tags string `json:"tags,omitempty" v-create:"isdefault"`
		Inputs   []CollectorConfigurationInput   `json:"inputs"`
		Outputs  []CollectorConfigurationOutput  `json:"outputs"`
		Snippets []CollectorConfigurationSnippet `json:"snippets"`
	}

	// CollectorConfigurationInput represents a Graylog's Collector Configuration Input.
	CollectorConfigurationInput struct {
		Backend    string                              `json:"backend"`
		Type       string                              `json:"type"`
		Name       string                              `json:"name"`
		InputID    string                              `json:"input_id"`
		ForwardTo  string                              `json:"forward_to"`
		Properties CollectorConfigurationInputProperty `json:"properties"`
	}

	// CollectorConfigurationInputProperty represents a Graylog's Collector Configuration Input properties.
	CollectorConfigurationInputProperty struct {
		Paths         string `json:"paths"`
		ExcludeFiles  string `json:"exclude_files"`
		ScanFrequency string `json:"scan_frequency"`
		Encoding      string `json:"encoding"`
		IgnoreOlder   string `json:"ignore_older"`
		DocumentType  string `json:"document_type"`
		ExcludeLines  string `json:"exclude_lines"`
		IncludeLines  string `json:"include_lines"`
		TailFiles     bool   `json:"tail_files"`
	}

	// CollectorConfigurationOutput represents a Graylog's Collector Configuration Output.
	CollectorConfigurationOutput struct {
		Backend    string                               `json:"backend"`
		Type       string                               `json:"type"`
		Name       string                               `json:"name"`
		OutputID   string                               `json:"output_id"`
		Properties CollectorConfigurationOutputProperty `json:"properties"`
	}

	// CollectorConfigurationOutputProperty represents a Graylog's Collector Configuration Output properties.
	CollectorConfigurationOutputProperty struct {
		Hosts string `json:"hosts"`
	}

	// CollectorConfigurationSnippet represents a Graylog's Collector Configuration Snippet.
	CollectorConfigurationSnippet struct {
		Backend   string `json:"backend"`
		Name      string `json:"name"`
		Snippet   string `json:"snippet"`
		SnippetID string `json:"snippet_id"`
	}
)

// CollectorConfigurationsBody represents Get Collector Configurations API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type CollectorConfigurationsBody struct {
	Configurations []CollectorConfiguration `json:"configurations"`
	Total          int                      `json:"total"`
}
