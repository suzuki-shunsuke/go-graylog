package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasCollectorConfigurationOutput returns whether the collector configuration exists.
func (lgc *Logic) HasCollectorConfigurationOutput(cfgID, outputID string) (bool, error) {
	return lgc.store.HasCollectorConfigurationOutput(cfgID, outputID)
}

// AddCollectorConfigurationOutput adds a collector configuration output to the mock server.
func (lgc *Logic) AddCollectorConfigurationOutput(id string, output *graylog.CollectorConfigurationOutput) (int, error) {
	if id == "" {
		return 400, fmt.Errorf("id is required")
	}
	if err := validator.CreateValidator.Struct(output); err != nil {
		return 400, err
	}
	if err := lgc.store.AddCollectorConfigurationOutput(id, output); err != nil {
		return 500, err
	}
	// 202 no content
	return 202, nil
}

// UpdateCollectorConfigurationOutput updates a collector configuration output.
func (lgc *Logic) UpdateCollectorConfigurationOutput(
	cfgID, outputID string, output *graylog.CollectorConfigurationOutput,
) (int, error) {
	if cfgID == "" {
		return 400, fmt.Errorf("collector configuration id is required")
	}
	if outputID == "" {
		return 400, fmt.Errorf("collector configuration output id is required")
	}
	if err := validator.UpdateValidator.Struct(output); err != nil {
		return 400, err
	}
	ok, err := lgc.HasCollectorConfigurationOutput(cfgID, outputID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the collector configuration output is not found")
	}

	if err := lgc.store.UpdateCollectorConfigurationOutput(cfgID, outputID, output); err != nil {
		return 500, err
	}
	// 202 no content
	return 200, nil
}

// DeleteCollectorConfigurationOutput deletes a collector configuration output from the mock server.
func (lgc *Logic) DeleteCollectorConfigurationOutput(cfgID, outputID string) (int, error) {
	ok, err := lgc.HasCollectorConfigurationOutput(cfgID, outputID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "configuration_id": cfgID,
			"output_id": outputID,
		}).Error("failed to check whether the collector configuration exists")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the collector configuration output is not found")
	}
	if err := lgc.store.DeleteCollectorConfigurationOutput(cfgID, outputID); err != nil {
		return 500, err
	}
	return 204, nil
}
