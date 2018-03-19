package graylog

// InputAttributes represents Input's attributes.
type InputAttributes struct {
	// OverrideSource string `json:"override_source,omitempty"`
	RecvBufferSize      int    `json:"recv_buffer_size,omitempty"`
	BindAddress         string `json:"bind_address,omitempty"`
	Port                int    `json:"port,omitempty"`
	DecompressSizeLimit int    `json:"decompress_size_limit,omitempty"`
}

// InputConfiguration represents Input's configuration.
type InputConfiguration struct {
	// ex. 0.0.0.0
	BindAddress string `json:"bind_address,omitempty" v-create:"required" v-update:"required"`
	Port        int    `json:"port,omitempty" v-create:"required" v-update:"required"`
	// ex. 262144
	RecvBufferSize int `json:"recv_buffer_size,omitempty" v-create:"required" v-update:"required"`
}

// Input represents Graylog Input.
type Input struct {
	// required
	Title         string              `json:"title,omitempty" v-create:"required" v-update:"required"`
	Type          string              `json:"type,omitempty" v-create:"required" v-update:"required"`
	Configuration *InputConfiguration `json:"configuration,omitempty" v-create:"required" v-update:"required"`

	// ex. "5a90d5c2c006c60001efc368"
	ID string `json:"id,omitempty" v-create:"isdefault" v-update:"required"`

	Global bool `json:"global,omitempty"`
	// ex. "2ad6b340-3e5f-4a96-ae81-040cfb8b6024"
	Node string `json:"node,omitempty"`
	// ex. 2018-02-24T03:02:26.001Z
	CreatedAt string `json:"created_at,omitempty" v-create:"isdefault" v-update:"isdefault"`
	// ex. "admin"
	CreatorUserID string           `json:"creator_user_id,omitempty" v-create:"isdefault" v-update:"isdefault"`
	Attributes    *InputAttributes `json:"attributes,omitempty" v-create:"isdefault"`
	// ContextPack `json:"context_pack,omitempty"`
	// StaticFields `json:"static_fields,omitempty"`
}

type InputsBody struct {
	Inputs []Input `json:"inputs"`
	Total  int     `json:"total"`
}
