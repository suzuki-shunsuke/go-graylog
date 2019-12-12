package graylog

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/suzuki-shunsuke/go-set"
)

const (
	// EmailAlarmCallbackType is a type of EmailAlarmCallback.
	EmailAlarmCallbackType = "org.graylog2.alarmcallbacks.EmailAlarmCallback"
	// HTTPAlarmCallbackType is a type of HTTPAlarmCallback.
	HTTPAlarmCallbackType = "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
	// SlackAlarmCallbackType is a type of SlackAlarmCallback.
	SlackAlarmCallbackType = "org.graylog2.plugins.slack.callback.SlackAlarmCallback"
)

type (
	// AlarmCallback represents an Alarm Callback.
	// https://docs.graylog.org/en/latest/pages/streams/alerts.html#alert-notifications-types-explained
	AlarmCallback struct {
		// ex. "org.graylog2.alarmcallbacks.HTTPAlarmCallback"
		ID            string                     `json:"id,omitempty" v-create:"isdefault"`
		StreamID      string                     `json:"stream_id,omitempty" v-create:"required"`
		Title         string                     `json:"title" v-create:"required"`
		CreatorUserID string                     `json:"creator_user_id,omitempty" v-create:"isdefault"`
		CreatedAt     string                     `json:"created_at,omitempty" v-create:"isdefault"`
		Configuration AlarmCallbackConfiguration `json:"configuration" v-create:"required"`
	}

	// AlarmCallbackConfiguration is an alarm callback configuration.
	AlarmCallbackConfiguration interface {
		AlarmCallbackType() string
	}

	// HTTPAlarmCallbackConfiguration represents a configuration of HTTPAlarmCallback.
	HTTPAlarmCallbackConfiguration struct {
		URL string `json:"url" v-create:"required"`
	}

	// EmailAlarmCallbackConfiguration represents a configuration of EmailAlarmCallback.
	EmailAlarmCallbackConfiguration struct {
		Sender         string     `json:"sender" v-create:"required"`
		Subject        string     `json:"subject" v-create:"required"`
		Body           string     `json:"body,omitempty"`
		UserReceivers  set.StrSet `json:"user_receivers,omitempty"`
		EmailReceivers set.StrSet `json:"email_receivers,omitempty"`
	}

	// SlackAlarmCallbackConfiguration represents a configuration of SlackAlarmCallback.
	// Note that SlackAlarmCallback is a third party plugin and not an official alarm callback.
	// https://github.com/graylog-labs/graylog-plugin-slack
	SlackAlarmCallbackConfiguration struct {
		Color         string `json:"color" v-create:"required"`
		WebhookURL    string `json:"webhook_url" v-create:"required"`
		Channel       string `json:"channel" v-create:"required"`
		IconURL       string `json:"icon_url,omitempty"`
		Graylog2URL   string `json:"graylog2_url,omitempty"`
		IconEmoji     string `json:"icon_emoji,omitempty"`
		UserName      string `json:"user_name,omitempty"`
		ProxyAddress  string `json:"proxy_address,omitempty"`
		CustomMessage string `json:"custom_message,omitempty"`
		BacklogItems  int    `json:"backlog_items,omitempty"`
		LinkNames     bool   `json:"link_names"`
		NotifyChannel bool   `json:"notify_channel"`
	}

	// AlarmCallbacksBody represents Get Alarm Callbacks API's response body.
	// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
	AlarmCallbacksBody struct {
		AlarmCallbacks []AlarmCallback `json:"alarmcallbacks"`
		Total          int             `json:"total"`
	}

	// GeneralAlarmCallbackConfiguration is a general third party's AlarmCallbackConfiguration.
	GeneralAlarmCallbackConfiguration struct {
		Type          string                 `json:"type"`
		Configuration map[string]interface{} `json:"configuration"`
	}
)

// Type returns an alarm callback type.
func (ac AlarmCallback) Type() string {
	if ac.Configuration == nil {
		return ""
	}
	return ac.Configuration.AlarmCallbackType()
}

// MarshalJSON returns JSON encoding of an AlarmCallback.
func (ac *AlarmCallback) MarshalJSON() ([]byte, error) {
	if ac == nil {
		return []byte("{}"), nil
	}
	type alias AlarmCallback
	return json.Marshal(struct {
		Type string `json:"type"`
		*alias
	}{
		Type:  ac.Type(),
		alias: (*alias)(ac),
	})
}

// AlarmCallbackType returns an alarm callback type.
func (ac *HTTPAlarmCallbackConfiguration) AlarmCallbackType() string {
	return HTTPAlarmCallbackType
}

// AlarmCallbackType returns an alarm callback type.
func (ac *EmailAlarmCallbackConfiguration) AlarmCallbackType() string {
	return EmailAlarmCallbackType
}

// AlarmCallbackType returns an alarm callback type.
func (ac *SlackAlarmCallbackConfiguration) AlarmCallbackType() string {
	return SlackAlarmCallbackType
}

// UnmarshalJSON unmarshals JSON into an AlarmCallback.
func (ac *AlarmCallback) UnmarshalJSON(b []byte) error {
	errMsg := "failed to unmarshal JSON to AlarmCallback"
	if ac == nil {
		return errors.New(errMsg + ": AlarmCallback is nil")
	}
	type alias AlarmCallback
	a := struct {
		Type          string          `json:"type"`
		Configuration json.RawMessage `json:"configuration"`
		*alias
	}{
		alias: (*alias)(ac),
	}
	if err := json.Unmarshal(b, &a); err != nil {
		return fmt.Errorf(errMsg+": %w", err)
	}
	switch a.Type {
	case EmailAlarmCallbackType:
		p := EmailAlarmCallbackConfiguration{}
		if err := json.Unmarshal(a.Configuration, &p); err != nil {
			return fmt.Errorf(errMsg+": %w", err)
		}
		ac.Configuration = &p
		return nil
	case HTTPAlarmCallbackType:
		p := HTTPAlarmCallbackConfiguration{}
		if err := json.Unmarshal(a.Configuration, &p); err != nil {
			return fmt.Errorf(errMsg+": %w", err)
		}
		ac.Configuration = &p
		return nil
	case SlackAlarmCallbackType:
		p := SlackAlarmCallbackConfiguration{}
		if err := json.Unmarshal(a.Configuration, &p); err != nil {
			return fmt.Errorf(errMsg+": %w", err)
		}
		ac.Configuration = &p
		return nil
	}
	p := map[string]interface{}{}
	if err := json.Unmarshal(a.Configuration, &p); err != nil {
		return fmt.Errorf(errMsg+": %w", err)
	}
	ac.Configuration = &GeneralAlarmCallbackConfiguration{
		Type: a.Type, Configuration: p,
	}
	return nil
}

// AlarmCallbackType returns an alarm callback type.
func (p *GeneralAlarmCallbackConfiguration) AlarmCallbackType() string {
	return p.Type
}

// MarshalJSON returns JSON encoding of GeneralAlarmCallbackConfiguration.
func (p *GeneralAlarmCallbackConfiguration) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(p.Configuration)
}
