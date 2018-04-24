package graylog

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/suzuki-shunsuke/go-graylog/util"
	"github.com/suzuki-shunsuke/go-ptr"
)

const (
	InputTypeAWSCloudTrail   string = "org.graylog.aws.inputs.cloudtrail.CloudTrailInput"
	InputTypeAWSFlowLogs     string = "org.graylog.aws.inputs.flowlogs.FlowLogsInput"
	InputTypeAWSLogs         string = "org.graylog.aws.inputs.cloudwatch.CloudWatchLogsInput"
	InputTypeBeats           string = "org.graylog.plugins.beats.BeatsInput"
	InputTypeCEFAMQP         string = "org.graylog.plugins.cef.input.CEFAmqpInput"
	InputTypeCEFKafka        string = "org.graylog.plugins.cef.input.CEFKafkaInput"
	InputTypeCEFTCP          string = "org.graylog.plugins.cef.input.CEFTCPInput"
	InputTypeCEFUDP          string = "org.graylog.plugins.cef.input.CEFUDPInput"
	InputTypeFakeHTTPMessage string = "org.graylog2.inputs.random.FakeHttpMessageInput"
	InputTypeGELFAMQP        string = "org.graylog2.inputs.gelf.amqp.GELFAMQPInput"
	InputTypeGELFHTTP        string = "org.graylog2.inputs.gelf.http.GELFHttpInput"
	InputTypeGELFKafka       string = "org.graylog2.inputs.gelf.kafka.GELFKafkaInput"
	InputTypeGELFTCP         string = "org.graylog2.inputs.gelf.tcp.GELFTCPInput"
	InputTypeGELFUDP         string = "org.graylog2.inputs.gelf.udp.GELFUDPInput"
	InputTypeJSONPath        string = "org.graylog2.inputs.misc.jsonpath.JsonPathInput"
	InputTypeNetFlowUDP      string = "org.graylog.plugins.netflow.inputs.NetFlowUdpInput"
	InputTypeRawAMQP         string = "org.graylog2.inputs.raw.amqp.RawAMQPInput"
	InputTypeSyslogAMQP      string = "org.graylog2.inputs.syslog.amqp.SyslogAMQPInput"
	InputTypeSyslogKafka     string = "org.graylog2.inputs.syslog.kafka.SyslogKafkaInput"
	InputTypeSyslogTCP       string = "org.graylog2.inputs.syslog.tcp.SyslogTCPInput"
	InputTypeSyslogUDP       string = "org.graylog2.inputs.syslog.udp.SyslogUDPInput"
)

var (
	// When update these fields variables, update also terraform graylog_input resource's document.
	InputAttributesIntFields = []string{
		"port", "recv_buffer_size", "heartbeat", "prefetch", "broker_port",
		"parallel_queues", "fetch_wait_max", "fetch_min_bytes", "threads",
		"max_message_size", "decompress_size_limit", "idle_writer_timeout",
		"max_chunk_size", "interval"}
	// When update these fields variables, update also terraform graylog_input resource's document.
	InputAttributesBoolFields = []string{
		"throttling_allowed", "tls_enable", "tcp_keepalive", "exchange_bind", "tls", "requeue_invalid_messages", "use_full_names", "use_null_delimiter", "enable_cors", "force_rdns", "store_full_message", "expand_structured_data", "allow_override_date"}
	// When update these fields variables, update also terraform graylog_input resource's document.
	InputAttributesStrFields = []string{
		"bind_address", "aws_region", "aws_assume_role_arn", "aws_access_key", "kinesis_stream_name", "aws_secret_key", "aws_sqs_region", "aws_s3_region", "aws_sqs_queue_name", "override_source", "tls_key_file", "tls_key_password", "tls_client_auth_cert_file", "tls_client_auth", "tls_cert_file", "timezone", "broker_vhost", "broker_username", "locale", "broker_password", "exchange", "routing_key", "broker_hostname", "queue", "topic_filter", "offset_reset", "zookeeper", "headers", "path", "target_url", "source", "timeunit", "netflow9_definitions_path"}
	inputAttrsList = []InputAttributes{
		&InputCloudTrailAttrs{},
		&InputAWSFlowLogsAttrs{},
		&InputAWSLogsAttrs{},
		&InputBeatsAttrs{},
		&InputCEFAMQPAttrs{},
		&InputCEFKafkaAttrs{},
		&InputCEFTCPAttrs{},
		&InputCEFUDPAttrs{},
		&InputFakeHTTPMessageAttrs{},
		&InputGELFAMQPAttrs{},
		&InputGELFHTTPAttrs{},
		&InputGELFKafkaAttrs{},
		&InputGELFTCPAttrs{},
		&InputGELFUDPAttrs{},
		&InputJSONPathAttrs{},
		&InputNetFlowUDPAttrs{},
		&InputRawAMQPAttrs{},
		&InputSyslogAMQPAttrs{},
		&InputSyslogKafkaAttrs{},
		&InputSyslogTCPAttrs{},
		&InputSyslogUDPAttrs{},
	}
	inputAttrsMap = map[string]InputAttributes{}
)

func init() {
	for _, attrs := range inputAttrsList {
		inputAttrsMap[attrs.InputType()] = attrs
	}
}

// NewInputAttrs returns a new InputAttributes.
func NewInputAttrs(t string) (InputAttributes, error) {
	a, ok := inputAttrsMap[t]
	if !ok {
		return &InputUnknownAttrs{inputType: t}, nil
	}
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("InputAttributes must be a pointer")
	}
	return reflect.New(reflect.Indirect(v).Type()).Interface().(InputAttributes), nil
}

// InputUnknownAttrs represents unknown type's Input Attributes.
type InputUnknownAttrs struct {
	inputType string
	Data      map[string]interface{}
}

// InputType is the implementation of the InputAttributes interface.
func (attrs InputUnknownAttrs) InputType() string {
	return attrs.inputType
}

// InputAttributes represents Input Attributes.
// A receiver must be a pointer.
type InputAttributes interface {
	InputType() string
}

// Input represents Graylog Input.
type Input struct {
	// required
	// Select a name of your new input that describes it.
	Title string `json:"title,omitempty" v-create:"required" v-update:"required"`
	Type  string `json:"type,omitempty" v-create:"required" v-update:"required"`
	// https://github.com/Graylog2/graylog2-server/issues/3480
	// update input overwrite attributes
	Attributes InputAttributes `json:"attributes,omitempty" v-create:"required" v-update:"required"`

	// ex. "5a90d5c2c006c60001efc368"
	ID string `json:"id,omitempty" v-create:"isdefault" v-update:"required,objectid"`

	// Should this input start on all nodes
	Global bool `json:"global,omitempty"`
	// On which node should this input start
	// ex. "2ad6b340-3e5f-4a96-ae81-040cfb8b6024"
	Node string `json:"node,omitempty"`
	// ex. 2018-02-24T03:02:26.001Z
	CreatedAt string `json:"created_at,omitempty" v-create:"isdefault"`
	// ex. "admin"
	CreatorUserID string `json:"creator_user_id,omitempty" v-create:"isdefault"`
	// ContextPack `json:"context_pack,omitempty"`
	// StaticFields `json:"static_fields,omitempty"`
}

// NewUpdateParams converts Input to InputUpdateParams.
func (input *Input) NewUpdateParams() *InputUpdateParams {
	return &InputUpdateParams{
		ID:         input.ID,
		Title:      input.Title,
		Type:       input.Type,
		Attributes: input.Attributes,
		Node:       input.Node,
		Global:     ptr.PBool(input.Global),
	}
}

// InputUpdateParams represents Graylog Input update API's parameter.
type InputUpdateParams struct {
	ID         string          `json:"id,omitempty" v-update:"required,objectid"`
	Title      string          `json:"title,omitempty" v-update:"required"`
	Type       string          `json:"type,omitempty" v-update:"required"`
	Attributes InputAttributes `json:"attributes,omitempty" v-update:"required"`
	Global     *bool           `json:"global,omitempty"`
	Node       string          `json:"node,omitempty"`
}

// InputUpdateParamsData represents InputUpdateParams's data.
// This is used for data conversion of InputUpdateParams.
// ex. json.Unmarshal
type InputUpdateParamsData struct {
	ID         string                 `json:"id,omitempty"`
	Title      string                 `json:"title,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Global     *bool                  `json:"global,omitempty"`
	Node       string                 `json:"node,omitempty"`
}

// InputData represents data of Input.
// This is used for data conversion of Input.
// ex. json.Unmarshal
type InputData struct {
	Title         string                 `json:"title,omitempty"`
	Type          string                 `json:"type,omitempty"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
	ID            string                 `json:"id,omitempty"`
	Global        bool                   `json:"global,omitempty"`
	Node          string                 `json:"node,omitempty"`
	CreatedAt     string                 `json:"created_at,omitempty"`
	CreatorUserID string                 `json:"creator_user_id,omitempty"`
}

// ToInputUpdateParams copies InputUpdateParamsData's data to InputUpdateParams.
func (d *InputUpdateParamsData) ToInputUpdateParams(input *InputUpdateParams) error {
	input.Title = d.Title
	input.Type = d.Type
	input.ID = d.ID
	input.Global = d.Global
	input.Node = d.Node
	attrs, err := NewInputAttrs(input.Type)
	if err != nil {
		return err
	}
	if _, ok := attrs.(*InputUnknownAttrs); ok {
		input.Attributes = InputUnknownAttrs{inputType: input.Type, Data: d.Attributes}
		return nil
	}
	if err := util.MSDecode(d.Attributes, attrs); err != nil {
		return err
	}
	input.Attributes = attrs
	return nil
}

// ToInput copies InputData's data to Input.
func (d *InputData) ToInput(input *Input) error {
	input.Title = d.Title
	input.Type = d.Type
	input.ID = d.ID
	input.Global = d.Global
	input.Node = d.Node
	input.CreatedAt = d.CreatedAt
	input.CreatorUserID = d.CreatorUserID
	attrs, err := NewInputAttrs(input.Type)
	if err != nil {
		return err
	}
	if _, ok := attrs.(*InputUnknownAttrs); ok {
		input.Attributes = InputUnknownAttrs{inputType: input.Type, Data: d.Attributes}
		return nil
	}
	if err := util.MSDecode(d.Attributes, attrs); err != nil {
		return err
	}
	input.Attributes = attrs
	return nil
}

// UnmarshalJSON is the implementation of the json.Unmarshaler interface.
func (input *Input) UnmarshalJSON(b []byte) error {
	d := &InputData{
		Title:         input.Title,
		Type:          input.Type,
		ID:            input.ID,
		Global:        input.Global,
		Node:          input.Node,
		CreatedAt:     input.CreatedAt,
		CreatorUserID: input.CreatorUserID,
		Attributes:    map[string]interface{}{},
	}
	if input.Attributes != nil {
		if err := util.MSDecode(input.Attributes, &(d.Attributes)); err != nil {
			return err
		}
	}
	if err := json.Unmarshal(b, d); err != nil {
		return err
	}
	return d.ToInput(input)
}

// InputsBody represents Get Inputs API's response body.
// Basically users don't use this struct, but this struct is public because some sub packages use this struct.
type InputsBody struct {
	Inputs []Input `json:"inputs"`
	Total  int     `json:"total"`
}
