package fastmvc

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/apex/log"
)

type Config interface {
	LoadConfig()
	ListenerAddr() string
	ParseLogLevel() log.Level
}

type BaseConfig struct {
	Host     string `default:""`
	Port     string `default:"8080"`
	LogLevel string `default:"info"`
}

func (c *BaseConfig) ListenerAddr() string {
	return c.Host + ":" + c.Port
}

func (c *BaseConfig) LoadConfig() {
	if err := envconfig.Process("", c); err != nil {
		log.Fatal(err.Error())
	}
}

func (c *BaseConfig) ParseLogLevel() log.Level {
	if level, err := log.ParseLevel(c.LogLevel); err != nil {
		return log.InfoLevel
	} else {
		return level
	}
}