package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/suzuki-shunsuke/go-set"
)

// validateRequestBody validates a request body and converts it into a map.
func validateRequestBody(
	b io.Reader, requiredFields, allowedFields, optionalFields *set.StrSet,
) (
	map[string]interface{}, int, error,
) {
	dec := json.NewDecoder(b)
	var a interface{}
	if err := dec.Decode(&a); err != nil {
		return nil, 400, fmt.Errorf(
			"failed to parse the request body as JSON: %s", err)
	}
	body, ok := a.(map[string]interface{})
	if !ok {
		return nil, 400, fmt.Errorf(
			"failed to parse the request body as a JSON object: %s", a)
	}
	if requiredFields != nil {
		for k := range requiredFields.ToMap(false) {
			if _, ok := body[k]; !ok {
				return body, 400, fmt.Errorf(
					`in the request body the field "%s" is required`, k)
			}
		}
	}
	if allowedFields != nil && allowedFields.Len() != 0 {
		allowedFields.AddSet(requiredFields)
		arr := make([]string, allowedFields.Len())
		i := 0
		for k := range allowedFields.ToMap(false) {
			arr[i] = k
			i++
		}
		for k := range body {
			if !allowedFields.Has(k) {
				return body, 400, fmt.Errorf(
					`in the request body an invalid field is found: "%s". The allowed fields: %s`,
					k, strings.Join(arr, ", "))
			}
		}
	}
	if optionalFields != nil && optionalFields.Len() != 0 {
		optionalFields.AddSets(requiredFields, allowedFields)
		for k := range body {
			if !optionalFields.Has(k) {
				delete(body, k)
			}
		}
	}
	return body, 200, nil
}
