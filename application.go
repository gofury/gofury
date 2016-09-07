package fastmvc

import (
	"fmt"
	"net"
	"runtime"
	"github.com/valyala/fasthttp"
	"github.com/apex/log"
)

type Application struct {
	name    string
	config  Config
	logger  *log.Logger
	handler fasthttp.RequestHandler
}

func NewApplication(name string, config Config, l *log.Logger, h fasthttp.RequestHandler) *Application {
	return &Application{
		name: name,
		config: config,
		logger: l,
		handler: h,
	}
}

func (app *Application) Start() {
	fmt.Printf("listening on %s\n", app.config.ListenerAddr())

	s := &fasthttp.Server{
		Handler: app.handler,
		Name:    app.name,
	}

	if err := s.Serve(app.getListener(app.config.ListenerAddr())); err != nil {
		app.logger.Fatalf("Error when serving incoming connections: %s", err)
	} else {
		app.logger.Info("started on")
	}
}

func (app *Application) getListener(listenAddr string) net.Listener {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ln, err := net.Listen("tcp4", listenAddr);
	if err != nil {
		app.logger.Fatal(err.Error())
	}
	return ln
}

