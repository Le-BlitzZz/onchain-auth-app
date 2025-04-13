package config

import (
	"path/filepath"
)

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

func (c *Config) AssetsPath() string {
	return c.options.AssetsPath
}

func (c *Config) StaticPath() string {
	return filepath.Join(c.AssetsPath(), "static")
}

func (c *Config) TemplatesPath() string {
	return filepath.Join(c.AssetsPath(), "templates")
}

func (c *Config) BuildPath() string {
	return filepath.Join(c.StaticPath(), "build")
}

func (c *Config) TemplateFiles() []string {
	results := make([]string, 0)

	matches, err := filepath.Glob(c.TemplatesPath() + "/[A-Za-z0-9]*.*")

	if err != nil {
		panic(err)
	}

	for _, tmpName := range matches {
		results = append(results, tmpName)
	}

	return results
}

func (c *Config) TemplateName() string {
	return "index.gohtml"
}
