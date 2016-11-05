package gofury

import (
	"testing"
	"github.com/valyala/fasthttp"
	"github.com/stretchr/testify/assert"
)

type DummyHealthCheck struct {
	name    string
	healthy bool
}

func (h *DummyHealthCheck) CheckHealth() bool {
	return h.healthy
}

func (h *DummyHealthCheck) HealthCheckName() string {
	return h.name
}

func TestCheckHealth(t *testing.T) {
	// given
	ctx := &fasthttp.RequestCtx{}

	// when
	CheckHealth(
		&DummyHealthCheck{name: "first", healthy: true},
		&DummyHealthCheck{name: "second", healthy: false},
	)(ctx)

	// then
	assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())
	assert.Equal(t, "application/json", string(ctx.Response.Header.ContentType()))

	healthCheckJSON := `
		{
			"first": true,
			"second": false

		}`
	assert.JSONEq(t, healthCheckJSON, string(ctx.Response.Body()))
}
