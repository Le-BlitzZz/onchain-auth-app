package config

type ClientConfig struct {
	ClientAssets *ClientAssets `json:"-"`
	JsUri        string        `json:"jsUri"`
	CssUri       string        `json:"cssUri"`
}

func (c *Config) GetClientConfig() *ClientConfig {
	a := c.ClientAssets()

	log.Infof("JsUri: %s", a.appJsUri())
	log.Infof("CssUri: %s", a.appCssUri())

	return &ClientConfig{
		ClientAssets: a,
		JsUri:        a.appJsUri(),
		CssUri:       a.appCssUri(),
	}
}
