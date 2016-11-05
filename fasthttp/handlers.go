package gofury

import (
	"github.com/gofury/fastjsonapi"
	"github.com/valyala/fasthttp"
	"github.com/gofury/gofury"
	"fmt"
)

func UnmarshalJSON(payload gofury.Model) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if len(ctx.PostBody()) == 0 {
			ctx.Error("Request body is empty", fasthttp.StatusUnprocessableEntity)
		}
		if err := payload.UnmarshalJSON(ctx.Request.Body()); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		}
	}
}

func UnmarshalJSONAPI(payload gofury.Model) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if len(ctx.PostBody()) == 0 {
			ctx.Error("Request body is empty", fasthttp.StatusUnprocessableEntity)
		}
		if err := fastjsonapi.UnmarshalPayload(ctx.PostBody(), payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		}
	}
}

func MarshalJSON(payload gofury.Model, status int) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		bytes, err := payload.MarshalJSON()
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
			return
		}
		ctx.Response.SetStatusCode(status)
		ctx.Write(bytes)

	}
}

func MarshalJSONAPI(payload gofury.Model, status int) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if err := fastjsonapi.MarshalOnePayload(ctx, payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		} else {
			ctx.Response.SetStatusCode(status)
		}
	}
}

func CheckHealth(healthChecks... gofury.HealthCheck) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")
		fmt.Fprintln(ctx, "{")
		for i, hc := range healthChecks {
			fmt.Fprintf(ctx, `	"%s": %t`, hc.HealthCheckName(), hc.CheckHealth())
			if i < len(healthChecks) - 1 {
				fmt.Fprintln(ctx, ",")
			} else {
				fmt.Fprintln(ctx)
			}
		}
		fmt.Fprintln(ctx, "}")
		ctx.SetStatusCode(fasthttp.StatusOK)
	}
}