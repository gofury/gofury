package fasthttp

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"runtime"
	"testing"
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

func TestHealthCheckSuccess(t *testing.T) {
	// given
	ctx := &fasthttp.RequestCtx{}

	// when
	CheckHealth(ctx,
		&DummyHealthCheck{name: "first", healthy: true},
		&DummyHealthCheck{name: "second", healthy: true},
	)

	// then
	assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())
	assert.Equal(t, "application/json", string(ctx.Response.Header.ContentType()))

	healthCheckJSON := "{\n"
	healthCheckJSON = healthCheckJSON + fmt.Sprintf("	\"number of goroutine\": %v,\n", runtime.NumGoroutine())
	healthCheckJSON = healthCheckJSON + `
			"first": true,
			"second": true

		}`
	assert.JSONEq(t, healthCheckJSON, string(ctx.Response.Body()))
}

func TestHealthCheckFail(t *testing.T) {
	// given
	ctx := &fasthttp.RequestCtx{}

	// when
	CheckHealth(ctx,
		&DummyHealthCheck{name: "first", healthy: true},
		&DummyHealthCheck{name: "second", healthy: false},
	)

	// then
	assert.Equal(t, fasthttp.StatusBadRequest, ctx.Response.StatusCode())
	assert.Equal(t, "application/json", string(ctx.Response.Header.ContentType()))

	healthCheckJSON := "{\n"
	healthCheckJSON = healthCheckJSON + fmt.Sprintf("	\"number of goroutine\": %v,\n", runtime.NumGoroutine())
	healthCheckJSON = healthCheckJSON + `
			"first": true,
			"second": false

		}`
	assert.JSONEq(t, healthCheckJSON, string(ctx.Response.Body()))
}

func TestHealthCheckWithNoHealthChecker(t *testing.T) {
	// given
	ctx := &fasthttp.RequestCtx{}

	// when
	CheckHealth()

	// then
	assert.Equal(t, fasthttp.StatusOK, ctx.Response.StatusCode())
	assert.Equal(t, "application/json", string(ctx.Response.Header.ContentType()))

	healthCheckJSON := "{\n"
	healthCheckJSON = healthCheckJSON + fmt.Sprintf("	\"number of goroutine\": %v\n", runtime.NumGoroutine())
	healthCheckJSON = healthCheckJSON + `

		}`
	assert.JSONEq(t, healthCheckJSON, string(ctx.Response.Body()))
}
