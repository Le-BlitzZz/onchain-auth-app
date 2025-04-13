package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ClientAssets struct {
	BuildPath                 string `json:"-"`
	BaseUri                   string `json:"-"`
	AppJs                     string `json:"app.js"`
	AppCss                    string `json:"app.css"`
	MaterialIconsRegularTtf   string `json:"MaterialIcons-Regular.ttf"`
	MaterialIconsRegularWoff  string `json:"MaterialIcons-Regular.woff"`
	MaterialIconsRegularEot   string `json:"MaterialIcons-Regular.eot"`
	MaterialIconsRegularWoff2 string `json:"MaterialIcons-Regular.woff2"`
}

// ClientAssets returns the frontend build assets.
func (c *Config) ClientAssets() *ClientAssets {
	result := newClientAssets(c.BuildPath(), StaticUri)

	if err := result.load("assets.json"); err != nil {
		log.Debugf("frontend: %s", err)
		log.Errorf("frontend: cannot read assets.json")
	}

	return result
}

// newClientAssets creates a new ClientAssets instance.
func newClientAssets(buildPath, baseUri string) *ClientAssets {
	return &ClientAssets{BuildPath: buildPath, BaseUri: baseUri}
}

func (a *ClientAssets) load(fileName string) error {
	jsonFile, err := os.ReadFile(filepath.Join(a.BuildPath, fileName))

	if err != nil {
		return err
	}

	return json.Unmarshal(jsonFile, a)
}

func (a *ClientAssets) appJsUri() string {
	if a.AppJs == "" {
		return ""
	}

	return fmt.Sprintf("%s/build/%s", a.BaseUri, a.AppJs)
}

func (a *ClientAssets) appCssUri() string {
	if a.AppCss == "" {
		return ""
	}

	return fmt.Sprintf("%s/build/%s", a.BaseUri, a.AppCss)
}
