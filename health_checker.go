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
	output, err := json.MarshalIndent(healthChecks, "", "  ")
	if err == nil {
		ctx.SetContentType("application/json")
		fmt.Fprintln(ctx, string(output))
	} else {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
	}
}
