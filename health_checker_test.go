package gofury

import (
	"testing"
	"github.com/valyala/fasthttp"
	"github.com/stretchr/testify/assert"
)

type DummyHealthCheck struct {
	Name    string
	Healthy bool
}

func (h *DummyHealthCheck) CheckHealth() bool {
	return h.Healthy
}

func (h *DummyHealthCheck) HealthCheckerName() string {
	return h.Name
}

func TestHealthCheck(t *testing.T) {
	// given
	healthChecks := HealthCheckers{
		&DummyHealthCheck{Name: "first", Healthy: true},
		&DummyHealthCheck{Name: "second", Healthy: false},
	}
	ctx := &fasthttp.RequestCtx{}

	// when
	HealthCheck(ctx, healthChecks)

	// then
	assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())
	assert.Equal(t, "application/json", string(ctx.Response.Header.ContentType()))

	healthCheckJSON := `[
		{
			"Name": "first",
			"Healthy": true
		},
		{
			"Name": "second",
			"Healthy": false
		}
	]`
	assert.JSONEq(t, healthCheckJSON, string(ctx.Response.Body()))
}
