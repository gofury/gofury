package gofury

import (
	"github.com/valyala/fasthttp"
	"fmt"
)

// HealthChecker interface.
type HealthChecker interface {
	Name() string
	CheckHealth() bool
}

func HealthCheck(ctx *fasthttp.RequestCtx, healthChecks... HealthChecker) {
	status := fasthttp.StatusOK
	ctx.SetContentType("application/json")
	fmt.Fprintln(ctx, "{")
	for i, hc := range healthChecks {
		result := hc.CheckHealth()
		fmt.Fprintf(ctx, `	"%s": %t`, hc.Name(), result)
		if (!result) {
			status = fasthttp.StatusBadRequest
		}
		if i < len(healthChecks)-1 {
			fmt.Fprintln(ctx, ",")
		} else {
			fmt.Fprintln(ctx)
		}
	}
	fmt.Fprintln(ctx, "}")
	ctx.SetStatusCode(status)
}
