package main

// Generate test data
// 1. call Graylog API by API browser
// 2. copy and paste the response to testdata/***.json
// 3. add type and update data
// 4. run this command and output test data
// 5. copy and paste the test data to testdata/***.go

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/sanity-io/litter"

	"github.com/suzuki-shunsuke/go-graylog/v11/graylog/graylog"
)

type (
	dumper interface {
		dump(string) error
	}
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
	}
}

func Main() error {
	options := make([]string, len(data))
	i := 0
	for k := range data {
		options[i] = k
		i++
	}
	file := ""
	prompt := &survey.Select{
		Message: "file",
		Options: options,
	}
	if err := survey.AskOne(prompt, &file); err != nil {
		return err
	}
	d, ok := data[file]
	if !ok {
		return errors.New("file is not found")
	}
	if err := d.dump("testdata/" + file + ".json"); err != nil {
		return err
	}
	return nil
}

func dump(input string, dest interface{}) error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(dest); err != nil {
		return err
	}
	options := litter.Options{
		HideZeroValues: true,
	}
	options.Dump(dest)
	return nil
}

var (
	data = map[string]dumper{
		"response_create_event_notification": EventNotification{},
		"request_create_event_notification":  EventNotification{},
		"event_definition/gets/response":     EventDefinitions{},
		"event_definition/get/response":      EventDefinition{},
		"event_definition/create/request":    EventDefinition{},
		"event_definition/create/response":   EventDefinition{},
		"event_definition/update/request":    EventDefinition{},
		"event_definition/update/response":   EventDefinition{},
		"event_notifications":                EventNotifications{},
		"users":                              Users{},
		"user":                               User{},
		"roles":                              Roles{},
		"role":                               Role{},
		"index_sets":                         IndexSets{},
		"index_set":                          IndexSet{},
		"inputs":                             Inputs{},
		"input":                              Input{},
		"streams":                            Streams{},
		"stream":                             Stream{},
		"stream_rules":                       StreamRules{},
		"stream_rule":                        StreamRule{},
		"dashboards":                         Dashboards{},
		"dashboard":                          Dashboard{},
		"stream_alarm_callbacks":             StreamAlarmCallbacks{},
		"slack_stream_alarm_callback":        StreamAlarmCallback{},
		"http_stream_alarm_callback":         StreamAlarmCallback{},
		"email_stream_alarm_callback":        StreamAlarmCallback{},
		"stream_alert_conditions":            StreamAlertConditions{},
		"outputs":                            Outputs{},
		"stdout_output":                      Output{},
		"views":                              Views{},
		"view":                               View{},
	}
)

type (
	Users struct {
		users graylog.UsersBody
	}

	User struct {
		user graylog.User
	}

	Roles struct {
		roles graylog.RolesBody
	}

	Role struct {
		role graylog.Role
	}

	IndexSets struct {
		data graylog.IndexSetsBody
	}

	IndexSet struct {
		data graylog.IndexSet
	}

	Inputs struct {
		data graylog.InputsBody
	}

	Input struct {
		data graylog.Input
	}

	Streams struct {
		data graylog.StreamsBody
	}

	Stream struct {
		data graylog.Stream
	}

	StreamRules struct {
		data graylog.StreamRulesBody
	}

	StreamRule struct {
		data graylog.StreamRule
	}

	Dashboards struct {
		data graylog.DashboardsBody
	}

	Dashboard struct {
		data graylog.Dashboard
	}

	StreamAlarmCallbacks struct {
		data graylog.AlarmCallbacksBody
	}

	StreamAlarmCallback struct {
		data graylog.AlarmCallback
	}

	StreamAlertConditions struct {
		data graylog.AlertConditionsBody
	}

	Output struct {
		data graylog.Output
	}

	Outputs struct {
		data graylog.OutputsBody
	}

	EventNotification struct {
		data graylog.EventNotification
	}

	EventNotifications struct {
		data graylog.EventNotificationsBody
	}

	EventDefinitions struct {
		data graylog.EventDefinitionsBody
	}

	EventDefinition struct {
		data graylog.EventDefinition
	}

	Views struct {
		data graylog.Views
	}

	View struct {
		data graylog.View
	}
)

func (users Users) dump(input string) error {
	return dump(input, &users.users)
}

func (user User) dump(input string) error {
	return dump(input, &user.user)
}

func (roles Roles) dump(input string) error {
	return dump(input, &roles.roles)
}

func (role Role) dump(input string) error {
	return dump(input, &role.role)
}

func (is IndexSets) dump(input string) error {
	return dump(input, &is.data)
}

func (is IndexSet) dump(input string) error {
	return dump(input, &is.data)
}

func (ip Inputs) dump(input string) error {
	return dump(input, &ip.data)
}

func (ip Input) dump(input string) error {
	return dump(input, &ip.data)
}

func (rule StreamRule) dump(input string) error {
	return dump(input, &rule.data)
}

func (rules StreamRules) dump(input string) error {
	return dump(input, &rules.data)
}

func (streams Streams) dump(input string) error {
	return dump(input, &streams.data)
}

func (stream Stream) dump(input string) error {
	return dump(input, &stream.data)
}

func (db Dashboard) dump(input string) error {
	return dump(input, &db.data)
}

func (db Dashboards) dump(input string) error {
	return dump(input, &db.data)
}

func (ac StreamAlarmCallbacks) dump(input string) error {
	return dump(input, &ac.data)
}

func (ac StreamAlarmCallback) dump(input string) error {
	return dump(input, &ac.data)
}

func (ac StreamAlertConditions) dump(input string) error {
	return dump(input, &ac.data)
}

func (output Output) dump(input string) error {
	return dump(input, &output.data)
}

func (outputs Outputs) dump(input string) error {
	return dump(input, &outputs.data)
}

func (notification EventNotification) dump(input string) error {
	return dump(input, &notification.data)
}

func (notifications EventNotifications) dump(input string) error {
	return dump(input, &notifications.data)
}

func (definitions EventDefinitions) dump(input string) error {
	return dump(input, &definitions.data)
}

func (definition EventDefinition) dump(input string) error {
	return dump(input, &definition.data)
}

func (v Views) dump(input string) error {
	return dump(input, &v.data)
}

func (v View) dump(input string) error {
	return dump(input, &v.data)
}
