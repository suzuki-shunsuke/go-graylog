package util

import (
	"github.com/mitchellh/mapstructure"
	"github.com/suzuki-shunsuke/go-set"
)

// MSDecode
func MSDecode(input, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil, Result: output, TagName: "json",
		DecodeHook: set.MapstructureDecodeHookFromListToStrSet,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(input)
}
