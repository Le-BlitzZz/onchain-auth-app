package config

func (c *Config) HttpHost() string {
	if c.options.HttpHost == "" {
		return "0.0.0.0"
	}

	return c.options.HttpHost
}

func (c *Config) HttpPort() int {
	if c.options.HttpPort == 0 {
		return 8080
	}

	return c.options.HttpPort
}
