package gofury

import (
	"testing"
	"github.com/apex/log"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestBaseConfig(t *testing.T) {
	// given
	os.Setenv("HOST", "api.amaysim.net/sms")
	os.Setenv("PORT", "443")
	os.Setenv("LOGLEVEL", "debug")

	// when
	c := BaseConfig{}
	c.LoadConfig()

	// then
	expected := BaseConfig {
			Host:"api.amaysim.net/sms",
			Port:"443",
			LogLevel:log.DebugLevel.String(),
		}

	assert.Equal(t, expected, c)
}
