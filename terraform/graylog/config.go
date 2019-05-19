package graylog

// Config represents terraform provider's configuration.
type Config struct {
	Endpoint     string
	AuthName     string
	AuthPassword string
	XRequestedBy string
	APIVersion   string
}

func (c *Config) loadAndValidate() error {
	return nil
}
