package graylog

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type (
	// AlertCondition represents an Alert Condition.
	// http://docs.graylog.org/en/2.4/pages/streams/alerts.html#conditions
	AlertCondition struct {
		ID            string                   `json:"id,omitempty"`
		CreatorUserID string                   `json:"creator_user_id,omitempty"`
		CreatedAt     string                   `json:"created_at,omitempty"`
		Title         string                   `json:"title" v-create:"required" v-update:"required"`
		InGrace       bool                     `json:"in_grace,omitempty"`
		Parameters    AlertConditionParameters `json:"parameters" v-create:"reqired" v-update:"required"`
	}

	// AlertConditionParameters represents Alert Condition's parameters.
	AlertConditionParameters interface {
		AlertConditionType() string
	}

	// FieldContentAlertConditionParameters represents Field Content Alert Condition's parameters.
	FieldContentAlertConditionParameters struct {
		Grace               int    `json:"grace"`
		Backlog             int    `json:"backlog"`
		RepeatNotifications bool   `json:"repeat_notifications,omitempty"`
		Field               string `json:"field,omitempty" v-create:"required"`
		Value               string `json:"value,omitempty" v-create:"required"`
		Query               string `json:"query,omitempty"`
	}

	// FieldAggregationAlertConditionParameters represents Field Aggregation Alert Condition's parameters.
	FieldAggregationAlertConditionParameters struct {
		Grace               int    `json:"grace"`
		Backlog             int    `json:"backlog"`
		Threshold           int    `json:"threshold"`
		Time                int    `json:"time" v-create:"required"`
		RepeatNotifications bool   `json:"repeat_notifications,omitempty"`
		Field               string `json:"field,omitempty" v-create:"required"`
		Query               string `json:"query,omitempty"`
		ThresholdType       string `json:"threshold_type,omitempty" v-create:"required"`
		Type                string `json:"type,omitempty" v-create:"required"`
	}

	// MessageCountAlertConditionParameters represents Field Aggregation Alert Condition's parameters.
	MessageCountAlertConditionParameters struct {
		Grace               int    `json:"grace"`
		Backlog             int    `json:"backlog"`
		Threshold           int    `json:"threshold"`
		Time                int    `json:"time"`
		RepeatNotifications bool   `json:"repeat_notifications,omitempty"`
		Query               string `json:"query,omitempty"`
		ThresholdType       string `json:"threshold_type,omitempty" v-create:"required"`
	}

	// AlertConditionsBody represents Get Alert Conditions API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	AlertConditionsBody struct {
		AlertConditions []AlertCondition `json:"conditions"`
		Total           int              `json:"total"`
	}

	// GeneralAlertConditionParameters is a general third party's alert condition parameters.
	GeneralAlertConditionParameters struct {
		Type       string
		Parameters map[string]interface{}
	}
)

// Type returns an alert condition type.
func (cond AlertCondition) Type() string {
	if cond.Parameters == nil {
		return ""
	}
	return cond.Parameters.AlertConditionType()
}

// MarshalJSON returns JSON encoding of an alert condition.
func (cond *AlertCondition) MarshalJSON() ([]byte, error) {
	if cond == nil {
		return []byte("{}"), nil
	}
	type alias AlertCondition
	return json.Marshal(struct {
		Type string `json:"type"`
		*alias
	}{
		Type:  cond.Type(),
		alias: (*alias)(cond),
	})
}

// UnmarshalJSON unmarshals JSON into an alert condition.
func (cond *AlertCondition) UnmarshalJSON(b []byte) error {
	errMsg := "failed to unmarshal JSON to alert condition"
	if cond == nil {
		return fmt.Errorf("%s: alert condition is nil", errMsg)
	}
	type alias AlertCondition
	a := struct {
		Type       string          `json:"type"`
		Parameters json.RawMessage `json:"parameters"`
		*alias
	}{
		alias: (*alias)(cond),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return errors.Wrap(err, errMsg)
	}
	switch a.Type {
	case "field_content_value":
		p := FieldContentAlertConditionParameters{}
		if err := json.Unmarshal(a.Parameters, &p); err != nil {
			return errors.Wrap(err, errMsg)
		}
		cond.Parameters = p
		return nil
	case "field_value":
		p := FieldAggregationAlertConditionParameters{}
		if err := json.Unmarshal(a.Parameters, &p); err != nil {
			return errors.Wrap(err, errMsg)
		}
		cond.Parameters = p
		return nil
	case "message_count":
		p := MessageCountAlertConditionParameters{}
		if err := json.Unmarshal(a.Parameters, &p); err != nil {
			return errors.Wrap(err, errMsg)
		}
		cond.Parameters = p
		return nil
	}
	p := map[string]interface{}{}
	if err := json.Unmarshal(a.Parameters, &p); err != nil {
		return errors.Wrap(err, errMsg)
	}
	cond.Parameters = GeneralAlertConditionParameters{
		Type: a.Type, Parameters: p,
	}
	return nil
}

// AlertConditionType returns an alert condition type.
func (p FieldContentAlertConditionParameters) AlertConditionType() string {
	return "field_content_value"
}

// AlertConditionType returns an alert condition type.
func (p FieldAggregationAlertConditionParameters) AlertConditionType() string {
	return "field_value"
}

// AlertConditionType returns an alert condition type.
func (p MessageCountAlertConditionParameters) AlertConditionType() string {
	return "message_count"
}

// AlertConditionType returns an alert condition type.
func (p GeneralAlertConditionParameters) AlertConditionType() string {
	return p.Type
}

// MarshalJSON returns JSON encoding of GeneralAlertConditionParameters.
func (p *GeneralAlertConditionParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Parameters)
}
