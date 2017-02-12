package fasthttp

import (
	"github.com/gofury/gofury"
	"github.com/buaazp/fasthttprouter"

	"github.com/apex/log"
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	app := gofury.BaseApplication{}

	// load configuration
	cfg := &gofury.HTTPConfig{}
	app.LoadConfig(cfg)

	// register services
	app.RegisterServices(createFastHTTPService(cfg, &log.Logger{}))

	// start up application
	app.StartUp()

	// shudown cleanly when application exits

}

func createFastHTTPService(cfg *gofury.HTTPConfig, logger *log.Logger) *gf_fasthttp.FastHTTPService {
	router := fasthttprouter.New()
	router.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprint(ctx, "pong")
	})

	return gf_fasthttp.NewFastHTTPService("Fast HTTP Server", cfg, logger, router.Handler)
}