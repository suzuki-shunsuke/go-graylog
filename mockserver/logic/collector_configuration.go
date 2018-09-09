package logic

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/validator"
)

// HasCollectorConfiguration returns whether the collector configuration exists.
func (lgc *Logic) HasCollectorConfiguration(id string) (bool, error) {
	return lgc.store.HasCollectorConfiguration(id)
}

// GetCollectorConfiguration returns a collector configuration.
// If a collector configuration is not found, returns an error.
func (lgc *Logic) GetCollectorConfiguration(id string) (*graylog.CollectorConfiguration, int, error) {
	if id == "" {
		return nil, 400, fmt.Errorf("collector configuration id is required")
	}
	if err := ValidateObjectID(id); err != nil {
		// unfortunately graylog returns not 400 but 404.
		return nil, 404, err
	}
	cfg, err := lgc.store.GetCollectorConfiguration(id)
	if err != nil {
		return cfg, 500, err
	}
	if cfg == nil {
		return nil, 404, fmt.Errorf("no collector configuration with id <%s> is found", id)
	}
	return cfg, 200, nil
}

// AddCollectorConfiguration adds a collector configuration to the mock server.
func (lgc *Logic) AddCollectorConfiguration(cfg *graylog.CollectorConfiguration) (int, error) {
	if err := validator.CreateValidator.Struct(cfg); err != nil {
		return 400, err
	}
	if err := lgc.store.AddCollectorConfiguration(cfg); err != nil {
		return 500, err
	}
	return 200, nil
}

// RenameCollectorConfiguration renames a collector configuration.
func (lgc *Logic) RenameCollectorConfiguration(id, name string) (*graylog.CollectorConfiguration, int, error) {
	if id == "" {
		return nil, 400, fmt.Errorf("id is required")
	}
	if name == "" {
		return nil, 400, fmt.Errorf("name is required")
	}
	ok, err := lgc.HasCollectorConfiguration(id)
	if err != nil {
		return nil, 500, err
	}
	if !ok {
		return nil, 404, fmt.Errorf("the collector configuration <%s> is not found", id)
	}

	cfg, err := lgc.store.RenameCollectorConfiguration(id, name)
	if err != nil {
		return nil, 500, err
	}
	return cfg, 200, nil
}

// DeleteCollectorConfiguration deletes a collector configuration from the mock server.
func (lgc *Logic) DeleteCollectorConfiguration(id string) (int, error) {
	ok, err := lgc.HasCollectorConfiguration(id)
	if err != nil {
		lgc.Logger().WithFields(log.Fields{
			"error": err, "id": id,
		}).Error("lgc.HasCollectorConfiguration() is failure")
		return 500, err
	}
	if !ok {
		return 404, fmt.Errorf("the collector configuration <%s> is not found", id)
	}
	if err := lgc.store.DeleteCollectorConfiguration(id); err != nil {
		return 500, err
	}
	return 200, nil
}

// GetCollectorConfigurations returns a list of collector configurations.
func (lgc *Logic) GetCollectorConfigurations() ([]graylog.CollectorConfiguration, int, int, error) {
	cfgs, total, err := lgc.store.GetCollectorConfigurations()
	if err != nil {
		return nil, 0, 500, err
	}
	return cfgs, total, 200, nil
}
