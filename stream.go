package graylog

import (
	"github.com/suzuki-shunsuke/go-ptr"
)

// CloneStream
// POST /streams/{streamID}/clone Clone a stream
// TestMatchStream
// POST /streams/{streamID}/testMatch Test matching of a stream against a supplied message

type (
	// Stream represents a steram.
	Stream struct {
		ID         string `json:"id,omitempty" v-create:"isdefault" v-update:"required,objectid"`
		Title      string `json:"title,omitempty" v-create:"required"`
		IndexSetID string `json:"index_set_id,omitempty" v-create:"required"`
		// ex. "2018-02-20T11:37:19.371Z"
		CreatedAt string `json:"created_at,omitempty" v-create:"isdefault"`
		// ex. local:admin
		CreatorUserID string `json:"creator_user_id,omitempty" v-create:"isdefault"`
		Description   string `json:"description,omitempty"`
		// ex. "AND"
		MatchingType                   string           `json:"matching_type,omitempty"`
		Outputs                        []Output         `json:"outputs,omitempty" v-create:"isdefault"`
		Rules                          []StreamRule     `json:"rules,omitempty"`
		AlertConditions                []AlertCondition `json:"alert_conditions,omitempty" v-create:"isdefault"`
		AlertReceivers                 *AlertReceivers  `json:"alert_receivers,omitempty" v-create:"isdefault"`
		Disabled                       bool             `json:"disabled,omitempty" v-create:"isdefault"`
		RemoveMatchesFromDefaultStream bool             `json:"remove_matches_from_default_stream,omitempty"`
		IsDefault                      bool             `json:"is_default,omitempty" v-create:"isdefault"`
		// ContentPack `json:"content_pack,omitempty"`
	}

	// StreamUpdateParams represents a steram update params.
	StreamUpdateParams struct {
		ID                             string           `json:"id,omitempty" v-update:"required,objectid"`
		Title                          string           `json:"title,omitempty"`
		IndexSetID                     string           `json:"index_set_id,omitempty"`
		Description                    string           `json:"description,omitempty"`
		Outputs                        []Output         `json:"outputs,omitempty"`
		MatchingType                   string           `json:"matching_type,omitempty"`
		Rules                          []StreamRule     `json:"rules,omitempty"`
		AlertConditions                []AlertCondition `json:"alert_conditions,omitempty"`
		AlertReceivers                 *AlertReceivers  `json:"alert_receivers,omitempty"`
		RemoveMatchesFromDefaultStream *bool            `json:"remove_matches_from_default_stream,omitempty"`
	}

	// Output represents an output.
	Output struct{}

	// AlertReceivers represents alert receivers.
	AlertReceivers struct {
		Emails []string `json:"emails,omitempty"`
		Users  []string `json:"users,omitempty"`
	}

	// StreamsBody represents Get Streams API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	StreamsBody struct {
		Total   int      `json:"total,omitempty"`
		Streams []Stream `json:"streams,omitempty"`
	}
)

// NewUpdateParams converts Stream to StreamUpdateParams.
func (stream *Stream) NewUpdateParams() *StreamUpdateParams {
	return &StreamUpdateParams{
		ID:                             stream.ID,
		Title:                          stream.Title,
		IndexSetID:                     stream.IndexSetID,
		Description:                    stream.Description,
		Outputs:                        stream.Outputs,
		MatchingType:                   stream.MatchingType,
		Rules:                          stream.Rules,
		AlertConditions:                stream.AlertConditions,
		AlertReceivers:                 stream.AlertReceivers,
		RemoveMatchesFromDefaultStream: ptr.PBool(stream.RemoveMatchesFromDefaultStream),
	}
}
