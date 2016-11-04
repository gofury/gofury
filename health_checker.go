package gofury

import (
	"github.com/valyala/fasthttp"
	"fmt"
	"encoding/json"
)

type HealthCheckers []HealthChecker

// HealthChecker interface.
type HealthChecker interface {
	CheckHealth() bool
	HealthCheckerName() string
}

func HealthCheck(ctx *fasthttp.RequestCtx, healthChecks HealthCheckers) {
	ctx.SetContentType("application/json")
	output, err := json.MarshalIndent(healthChecks, "", "    ")
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	fmt.Fprintln(ctx, string(output))
}
