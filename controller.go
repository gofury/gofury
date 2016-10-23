package gofury

import (
	"github.com/gofury/fastjsonapi"
	"github.com/valyala/fasthttp"
)

type Controller interface {
	UnmarshalRequest(m Model) fasthttp.RequestHandler
	MarshalResponse(m Model, status int) fasthttp.RequestHandler
}

type BaseController struct {
}

func (c *BaseController) UnmarshalRequest(payload Model) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if len(ctx.PostBody()) == 0 {
			ctx.Error("Request body is empty", fasthttp.StatusUnprocessableEntity)
		} else if err := fastjsonapi.UnmarshalPayload(ctx.PostBody(), payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		}
	}
}

func (c *BaseController) MarshalResponse(payload Model, status int) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if err := fastjsonapi.MarshalOnePayload(ctx, payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		} else {
			ctx.Response.SetStatusCode(status)
		}
	}
}
