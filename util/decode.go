package util

import (
	"github.com/mitchellh/mapstructure"
	"github.com/suzuki-shunsuke/go-set"
)

// MSDecode assigns input's data to outputs with mapstructure.
// output must be a pointer to a map or struct.
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
