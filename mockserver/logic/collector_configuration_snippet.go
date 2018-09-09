package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasCollectorConfigurationSnippet returns whether the collector configuration exists.
func (lgc *Logic) HasCollectorConfigurationSnippet(cfgID, snippetID string) (bool, error) {
	return lgc.store.HasCollectorConfigurationSnippet(cfgID, snippetID)
}

// AddCollectorConfigurationSnippet adds a collector configuration snippet to the mock server.
func (lgc *Logic) AddCollectorConfigurationSnippet(id string, snippet *graylog.CollectorConfigurationSnippet) (int, error) {
	if id == "" {
		return 400, fmt.Errorf("id is required")
	}
	if err := validator.CreateValidator.Struct(snippet); err != nil {
		return 400, err
	}
	if err := lgc.store.AddCollectorConfigurationSnippet(id, snippet); err != nil {
		return 500, err
	}
	// 202 no content
	return 202, nil
}

// UpdateCollectorConfigurationSnippet updates a collector configuration snippet.
func (lgc *Logic) UpdateCollectorConfigurationSnippet(
	cfgID, snippetID string, snippet *graylog.CollectorConfigurationSnippet,
) (int, error) {
	if cfgID == "" {
		return 400, fmt.Errorf("collector configuration id is required")
	}
	if snippetID == "" {
		return 400, fmt.Errorf("collector configuration snippet id is required")
	}
	if err := validator.UpdateValidator.Struct(snippet); err != nil {
		return 400, err
	}
	ok, err := lgc.HasCollectorConfigurationSnippet(cfgID, snippetID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the collector configuration snippet is not found")
	}

	if err := lgc.store.UpdateCollectorConfigurationSnippet(cfgID, snippetID, snippet); err != nil {
		return 500, err
	}
	// 202 no content
	return 200, nil
}

// DeleteCollectorConfigurationSnippet deletes a collector configuration snippet from the mock server.
func (lgc *Logic) DeleteCollectorConfigurationSnippet(cfgID, snippetID string) (int, error) {
	ok, err := lgc.HasCollectorConfigurationSnippet(cfgID, snippetID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "configuration_id": cfgID,
			"snippet_id": snippetID,
		}).Error("failed to check whether the collector configuration exists")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the collector configuration snippet is not found")
	}
	if err := lgc.store.DeleteCollectorConfigurationSnippet(cfgID, snippetID); err != nil {
		return 500, err
	}
	return 204, nil
}
