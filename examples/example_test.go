package examples

import (
	"github.com/gofury/gofury"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	gofury_fasthttp "github.com/gofury/gofury/fasthttp"

	"github.com/apex/log"
	"fmt"
)

func main() {
	app := gofury.BaseApplication{}

	// load configuration
	cfg := &gofury.HTTPConfig{}
	app.LoadConfig(cfg)

	// register services
	app.RegisterServices(createFastHTTPService(cfg, &log.Logger{}), createQueueService())

	// start up application
	app.StartUp()

	// shudown cleanly when application exits
	defer app.ShutDown()
}

func createFastHTTPService(cfg *gofury.HTTPConfig, logger *log.Logger) *gofury_fasthttp.FastHTTPService {
	router := fasthttprouter.New()
	router.GET("/ping", func(ctx *fasthttp.RequestCtx) {
		fmt.Fprint(ctx, "pong")
	})

	return gofury_fasthttp.NewFastHTTPService("Fast HTTP Server", cfg, logger, router.Handler)
}

func createQueueService() gofury.QueueService {
	return nil
}