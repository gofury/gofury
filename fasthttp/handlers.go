package fasthttp

import (
	"github.com/gofury/fastjsonapi"
	"github.com/valyala/fasthttp"
	"github.com/gofury/gofury"
	"fmt"
	"runtime"
	"encoding/json"
)

func JSONUnmarshaller(payload json.Unmarshaler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if len(ctx.PostBody()) == 0 {
			ctx.Error("Request body is empty", fasthttp.StatusUnprocessableEntity)
		}
		if err := payload.UnmarshalJSON(ctx.Request.Body()); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		}
	}
}

func JSONAPIUnmarshaller(payload json.Unmarshaler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if len(ctx.PostBody()) == 0 {
			ctx.Error("Request body is empty", fasthttp.StatusUnprocessableEntity)
		}
		if err := fastjsonapi.UnmarshalPayload(ctx.PostBody(), payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		}
	}
}

func JSONMarshaller(payload json.Marshaler, status int) fasthttp.RequestHandler {
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

func JSONAPIMarshaller(payload json.Marshaler, status int) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if err := fastjsonapi.MarshalOnePayload(ctx, payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		} else {
			ctx.Response.SetStatusCode(status)
		}
	}
}


func HealthCheckHandler(healthChecks ...gofury.HealthCheck) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		status := fasthttp.StatusOK
		ctx.SetContentType("application/json")
		fmt.Fprintln(ctx, "{")
		fmt.Fprintf(ctx, `	"number of goroutine": %v`, runtime.NumGoroutine())
		if len(healthChecks) > 0 {
			fmt.Fprintln(ctx, ",")
		}
		for i, hc := range healthChecks {
			result := hc.CheckHealth()
			fmt.Fprintf(ctx, `	"%s": %t`, hc.HealthCheckName(), result)
			if !result {
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

}
