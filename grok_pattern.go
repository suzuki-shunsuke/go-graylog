package graylog

type (
	// GrokPattern represents a grok pattern.
	GrokPattern struct {
		ID      string `json:"id,omitempty"`
		Name    string `json:"name"`
		Pattern string `json:"pattern"`
		// ContentPack interface{} `json:"content_pack"`
	}

	// GrokPatternsBody represents Get Grok Pattern API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	GrokPatternsBody struct {
		Patterns []GrokPattern `json:"patterns"`
	}
)
