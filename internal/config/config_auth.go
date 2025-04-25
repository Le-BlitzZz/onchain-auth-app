package config

func (c *Config) DefaultUser() string {
	return c.options.DefaultUser
}

func (c *Config) DefaultUserPassword() string {
	return c.options.DefaultPassword
}
