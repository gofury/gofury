package gofury

import (
	"testing"
	"github.com/apex/log"
	"os"
	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	*Logconfig
	*HTTPConfig
}

func TestLoadConfig(t *testing.T) {
	// given
	os.Setenv("HOST", "api.gofury.com/furiousandfast")
	os.Setenv("PORT", "443")
	os.Setenv("LOGLEVEL", "warn")
	app := BaseApplication{}

	// when
	c := &TestConfig{}
	app.LoadConfig(c)

	// then
	expected := &TestConfig{
		&Logconfig{
			LogLevel:log.WarnLevel.String(),
		},
		&HTTPConfig{
			Host:"api.gofury.com/furiousandfast",
			Port:"443",
		},
	}

	assert.Equal(t, expected, c)
}
