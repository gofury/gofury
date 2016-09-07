package fastmvc

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/apex/log"
)

type Config interface {
	LoadConfig()
	ListenerAddr() string
}

type BaseConfig struct {
	Host     string `default:""`
	Port     string `default:"8080"`
	LogLevel log.Level `default:"INFO"`
}

func (c *BaseConfig) ListenerAddr() string {
	return c.Host + ":" + c.Port
}

func (c *BaseConfig) LoadConfig() {
	if err := envconfig.Process("", c); err != nil {
		log.Fatal(err.Error())
	}
}