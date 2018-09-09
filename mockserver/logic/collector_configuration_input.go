package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasCollectorConfigurationInput returns whether the collector configuration exists.
func (lgc *Logic) HasCollectorConfigurationInput(cfgID, inputID string) (bool, error) {
	return lgc.store.HasCollectorConfigurationInput(cfgID, inputID)
}

// AddCollectorConfigurationInput adds a collector configuration input to the mock server.
func (lgc *Logic) AddCollectorConfigurationInput(id string, input *graylog.CollectorConfigurationInput) (int, error) {
	if id == "" {
		return 400, fmt.Errorf("id is required")
	}
	if err := validator.CreateValidator.Struct(input); err != nil {
		return 400, err
	}
	if err := lgc.store.AddCollectorConfigurationInput(id, input); err != nil {
		return 500, err
	}
	// 202 no content
	return 202, nil
}

// UpdateCollectorConfigurationInput updates a collector configuration input.
func (lgc *Logic) UpdateCollectorConfigurationInput(
	cfgID, inputID string, input *graylog.CollectorConfigurationInput,
) (int, error) {
	if cfgID == "" {
		return 400, fmt.Errorf("collector configuration id is required")
	}
	if inputID == "" {
		return 400, fmt.Errorf("collector configuration input id is required")
	}
	if err := validator.UpdateValidator.Struct(input); err != nil {
		return 400, err
	}
	ok, err := lgc.HasCollectorConfigurationInput(cfgID, inputID)
	if err != nil {
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the collector configuration input is not found")
	}

	if err := lgc.store.UpdateCollectorConfigurationInput(cfgID, inputID, input); err != nil {
		return 500, err
	}
	// 202 no content
	return 200, nil
}

// DeleteCollectorConfigurationInput deletes a collector configuration input from the mock server.
func (lgc *Logic) DeleteCollectorConfigurationInput(cfgID, inputID string) (int, error) {
	ok, err := lgc.HasCollectorConfigurationInput(cfgID, inputID)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "configuration_id": cfgID,
			"input_id": inputID,
		}).Error("failed to check whether the collector configuration exists")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the collector configuration input is not found")
	}
	if err := lgc.store.DeleteCollectorConfigurationInput(cfgID, inputID); err != nil {
		return 500, err
	}
	return 204, nil
}
