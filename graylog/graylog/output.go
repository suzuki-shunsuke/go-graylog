package graylog

type (
	// Output represents an output.
	Output struct {
		ID            string      `json:"id,omitempty"`
		Title         string      `json:"title"`
		Type          string      `json:"type"`
		CreatorUserID string      `json:"creator_user_id,omitempty"`
		CreatedAt     string      `json:"created_at,omitempty"`
		Configuration interface{} `json:"configuration"`
		// ContentPack interface{} `json:"content_pack"`
	}

	// OutputsBody represents Get Outputs API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	OutputsBody struct {
		Outputs []Output `json:"outputs"`
		Total   int      `json:"total"`
	}
)
