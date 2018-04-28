package graylog

import (
	"fmt"

	"github.com/suzuki-shunsuke/go-graylog/util"
)

// InputUpdateParamsData represents InputUpdateParams's data.
// This is used for data conversion of InputUpdateParams.
// ex. json.Unmarshal
type InputUpdateParamsData struct {
	ID         string                 `json:"id,omitempty"`
	Title      string                 `json:"title,omitempty"`
	Type       string                 `json:"type,omitempty"`
	Node       string                 `json:"node,omitempty"`
	Global     *bool                  `json:"global,omitempty"`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// InputData represents data of Input.
// This is used for data conversion of Input.
// ex. json.Unmarshal
type InputData struct {
	Title         string                 `json:"title,omitempty"`
	Type          string                 `json:"type,omitempty"`
	ID            string                 `json:"id,omitempty"`
	Node          string                 `json:"node,omitempty"`
	CreatedAt     string                 `json:"created_at,omitempty"`
	CreatorUserID string                 `json:"creator_user_id,omitempty"`
	Global        bool                   `json:"global,omitempty"`
	Attributes    map[string]interface{} `json:"attributes,omitempty"`
}

// ToInputUpdateParams copies InputUpdateParamsData's data to InputUpdateParams.
func (d *InputUpdateParamsData) ToInputUpdateParams(input *InputUpdateParams) error {
	input.Title = d.Title
	input.Type = d.Type
	input.ID = d.ID
	input.Global = d.Global
	input.Node = d.Node
	attrs := NewInputAttrs(input.Type)
	if _, ok := attrs.(*InputUnknownAttrs); ok {
		input.Attributes = &InputUnknownAttrs{inputType: input.Type, Data: d.Attributes}
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
	if input.Type() != "" && input.Type() != d.Type {
		return fmt.Errorf("input type is different")
	}
	if input.Attributes != nil && input.Attributes.InputType() != d.Type {
		return fmt.Errorf("input type is different")
	}
	input.Title = d.Title
	input.ID = d.ID
	input.Global = d.Global
	input.Node = d.Node
	input.CreatedAt = d.CreatedAt
	input.CreatorUserID = d.CreatorUserID
	attrs := NewInputAttrs(d.Type)
	if _, ok := attrs.(*InputUnknownAttrs); ok {
		input.Attributes = &InputUnknownAttrs{inputType: input.Type(), Data: d.Attributes}
		return nil
	}
	if err := util.MSDecode(d.Attributes, attrs); err != nil {
		return err
	}
	input.Attributes = attrs
	return nil
}
