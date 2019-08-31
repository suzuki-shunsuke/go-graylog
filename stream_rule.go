package graylog

import (
	"github.com/suzuki-shunsuke/go-ptr"
)

type (
	// StreamRule represents a stream rule.
	StreamRule struct {
		ID          string `json:"id,omitempty" v-create:"isdefault"`
		StreamID    string `json:"stream_id,omitempty" v-create:"required"`
		Field       string `json:"field,omitempty" v-create:"required"`
		Value       string `json:"value,omitempty" v-create:"required"`
		Description string `json:"description,omitempty"`
		Type        int    `json:"type,omitempty"`
		Inverted    bool   `json:"inverted"`
	}

	// StreamRuleUpdateParams represents Update Stream API's paramteres.
	StreamRuleUpdateParams struct {
		ID          string `json:"id,omitempty" v-update:"required,objectid"`
		StreamID    string `json:"stream_id,omitempty" v-update:"required,objectid"`
		Field       string `json:"field,omitempty" v-update:"required"`
		Value       string `json:"value,omitempty" v-update:"required"`
		Description string `json:"description,omitempty"`
		Type        *int   `json:"type,omitempty"`
		Inverted    *bool  `json:"inverted,omitempty"`
	}

	// StreamRulesBody represents Get stream rules API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	StreamRulesBody struct {
		Total       int          `json:"total"`
		StreamRules []StreamRule `json:"stream_rules"`
	}
)

// NewUpdateParams converts StreamRule to StreamRuleUpdateParams.
func (rule *StreamRule) NewUpdateParams() *StreamRuleUpdateParams {
	return &StreamRuleUpdateParams{
		ID:          rule.ID,
		StreamID:    rule.StreamID,
		Field:       rule.Field,
		Description: rule.Description,
		Type:        ptr.PInt(rule.Type),
		Inverted:    ptr.PBool(rule.Inverted),
		Value:       rule.Value,
	}
}
