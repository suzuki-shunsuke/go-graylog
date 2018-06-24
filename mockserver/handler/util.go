package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/suzuki-shunsuke/go-set"
)

type validateReqBodyPrms struct {
	Required     set.StrSet
	Optional     set.StrSet
	Ignored      set.StrSet
	Forbidden    set.StrSet
	ExtForbidden bool
}

// validateRequestBody validates a request body and converts it into a map.
func validateRequestBody(b io.Reader, prms *validateReqBodyPrms) (map[string]interface{}, int, error) {
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
	if prms.Required != nil {
		for k := range prms.Required.ToMap(false) {
			if _, ok := body[k]; !ok {
				return body, 400, fmt.Errorf(
					`in the request body the field "%s" is required`, k)
			}
		}
	}
	allowedFields := set.NewStrSet()
	allowedFields.AddSets(prms.Required, prms.Optional, prms.Ignored)
	for k := range body {
		if prms.Required != nil && prms.Required.Has(k) {
			continue
		}
		if prms.Optional != nil && prms.Optional.Has(k) {
			continue
		}
		if prms.Ignored != nil && prms.Ignored.Has(k) {
			delete(body, k)
			continue
		}
		if prms.Forbidden != nil && prms.Forbidden.Has(k) || prms.ExtForbidden {
			return body, 400, fmt.Errorf(
				`in the request body an invalid field is found: "%s". The allowed fields: %s`,
				k, strings.Join(allowedFields.ToList(), ", "))
		}
	}
	return body, 200, nil
}
