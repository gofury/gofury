package gofury

type HTTPConfig struct {
	Host     string `default:""`
	Port     string `default:"8080"`
}

func (c *HTTPConfig) ListenerAddr() string {
	return c.Host + ":" + c.Port
}