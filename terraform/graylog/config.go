package graylog

type Config struct {
	Endpoint     string
	AuthName     string
	AuthPassword string
}

func (c *Config) loadAndValidate() error {
	return nil
}
