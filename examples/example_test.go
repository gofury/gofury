package examples

import (
	"github.com/gofury/gofury"
	"github.com/buaazp/fasthttprouter"
	"fmt"

	fasthttpservice "github.com/gofury/gofury/fasthttp"

	"github.com/apex/log"
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

func createFastHTTPService(cfg *gofury.HTTPConfig, logger *log.Logger) *fasthttpservice.FastHTTPService {
	router := fasthttprouter.New()
	router.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprint(ctx, "pong")
	})

	return fasthttpservice.NewFastHTTPService("Fast HTTP Server", cfg, logger, router.Handler)
}