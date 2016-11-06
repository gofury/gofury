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

func (h *DummyHealthCheck) Name() string {
	return h.name
}

func TestHealthCheckSuccess(t *testing.T) {
	// given
	ctx := &fasthttp.RequestCtx{}

	// when
	HealthCheck(ctx,
		&DummyHealthCheck{name: "first", healthy: true},
		&DummyHealthCheck{name: "second", healthy: true},
	)

	// then
	assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())
	assert.Equal(t, "application/json", string(ctx.Response.Header.ContentType()))

	healthCheckJSON := `
		{
			"first": true,
			"second": true

		}`
	assert.JSONEq(t, healthCheckJSON, string(ctx.Response.Body()))
}

func TestHealthCheckFail(t *testing.T) {
	// given
	ctx := &fasthttp.RequestCtx{}

	// when
	HealthCheck(ctx,
		&DummyHealthCheck{name: "first", healthy: true},
		&DummyHealthCheck{name: "second", healthy: false},
	)

	// then
	assert.Equal(t, fasthttp.StatusBadRequest, ctx.Response.StatusCode())
	assert.Equal(t, "application/json", string(ctx.Response.Header.ContentType()))

	healthCheckJSON := `
		{
			"first": true,
			"second": false

		}`
	assert.JSONEq(t, healthCheckJSON, string(ctx.Response.Body()))
}
