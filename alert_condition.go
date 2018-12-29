package graylog

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// AlertCondition represents an Alert Condition.
// http://docs.graylog.org/en/2.4/pages/streams/alerts.html#conditions
type AlertCondition struct {
	ID            string                   `json:"id"`
	CreatorUserID string                   `json:"creator_user_id"`
	CreatedAt     string                   `json:"created_at"`
	Title         string                   `json:"title" v-create:"required" v-update:"required"`
	InGrace       bool                     `json:"in_grace"`
	Parameters    AlertConditionParameters `json:"parameters" v-create:"reqired" v-update:"required"`
}

func (cond AlertCondition) Type() string {
	if cond.Parameters == nil {
		return ""
	}
	return cond.Parameters.AlertConditionType()
}

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
	return nil
}

// AlertConditionParameters represents Alert Condition's parameters.
type AlertConditionParameters interface {
	AlertConditionType() string
}

// FieldContentAlertConditionParameters represents Field Content Alert Condition's parameters.
type FieldContentAlertConditionParameters struct {
	Grace               int    `json:"grace" v-create:"required"`
	Backlog             int    `json:"backlog" v-create:"required"`
	RepeatNotifications bool   `json:"repeat_notifications"`
	Field               string `json:"field" v-create:"required"`
	Value               string `json:"value" v-create:"required"`
	Query               string `json:"query"`
}

func (p FieldContentAlertConditionParameters) AlertConditionType() string {
	return "field_content_value"
}

// FieldAggregationAlertConditionParameters represents Field Aggregation Alert Condition's parameters.
type FieldAggregationAlertConditionParameters struct {
	Grace               int    `json:"grace" v-create:"required"`
	Backlog             int    `json:"backlog" v-create:"required"`
	Threshold           int    `json:"threshold"`
	Time                int    `json:"time"`
	RepeatNotifications bool   `json:"repeat_notifications"`
	Field               string `json:"field" v-create:"required"`
	Query               string `json:"query"`
	ThresholdType       string `json:"threshold_type"`
}

func (p FieldAggregationAlertConditionParameters) AlertConditionType() string {
	return "field_value"
}

// MessageCountAlertConditionParameters represents Field Aggregation Alert Condition's parameters.
type MessageCountAlertConditionParameters struct {
	Grace               int    `json:"grace" v-create:"required"`
	Backlog             int    `json:"backlog" v-create:"required"`
	Threshold           int    `json:"threshold"`
	Time                int    `json:"time"`
	RepeatNotifications bool   `json:"repeat_notifications"`
	Query               string `json:"query"`
	ThresholdType       string `json:"threshold_type"`
}

func (p MessageCountAlertConditionParameters) AlertConditionType() string {
	return "message_count"
}

// AlertConditionsBody represents Get Alert Conditions API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type AlertConditionsBody struct {
	AlertConditions []AlertCondition `json:"conditions"`
	Total           int              `json:"total"`
}
