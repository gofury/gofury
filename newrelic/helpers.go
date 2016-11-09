package newrelic

import "github.com/newrelic/go-agent"

// NullNewRelic returns a disabled New Relic appliaction.
func NullNewRelic() newrelic.Application {
	config := newrelic.NewConfig("application", "")
	config.Enabled = false
	app, _ := newrelic.NewApplication(config)

	return app
}

