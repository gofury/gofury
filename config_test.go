package gofury

import (
	"testing"
	"github.com/apex/log"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestBaseConfig(t *testing.T) {
	// given
	os.Setenv("HOST", "api.gofury.com/furiousandfast")
	os.Setenv("PORT", "443")
	os.Setenv("LOGLEVEL", "warn")

	// when
	c := BaseConfig{}
	c.LoadConfig()

	// then
	expected := BaseConfig {
			Host:"api.gofury.com/furiousandfast",
			Port:"443",
		LogLevel:log.WarnLevel.String(),
		}

	assert.Equal(t, expected, c)
}
