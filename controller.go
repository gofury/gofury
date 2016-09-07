package fastmvc

import (
	"github.com/valyala/fasthttp"
	"code.amaysim.net/vulcan/sms-gateway/models"
	"github.com/gofury/fastjsonapi"
)

type Controller interface {
	UnmarshalRequest(m models.Model) fasthttp.RequestHandler
	MarshalResponse(m models.Model, status int) fasthttp.RequestHandler
}

type BaseController struct {

}

func (c *BaseController) UnmarshalRequest(payload models.Model) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if (len(ctx.PostBody()) == 0) {
			ctx.Error("Request body is empty", fasthttp.StatusUnprocessableEntity)
		} else if err := fastjsonapi.UnmarshalPayload(ctx.PostBody(), payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		} else {
			ctx.SetUserValue("payload", payload)
		}
	}
}

func (c *BaseController) MarshalResponse(payload models.Model, status int) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if err := fastjsonapi.MarshalOnePayload(ctx, payload); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusUnprocessableEntity)
		} else {
			ctx.Response.SetStatusCode(status)
		}
	}
}
