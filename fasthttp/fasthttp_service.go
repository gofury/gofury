package fasthttp

import (
	"net"
	"runtime"
	"github.com/valyala/fasthttp"
	"github.com/apex/log"
	"github.com/gofury/gofury"
)

type FastHTTPService struct {
	name     string
	config   *gofury.HTTPConfig
	logger   *log.Logger
	handler  fasthttp.RequestHandler
	listener net.Listener
}

func NewFastHTTPService(name string, config *gofury.HTTPConfig, l *log.Logger, h fasthttp.RequestHandler) *FastHTTPService {
	return &FastHTTPService{
		name: name,
		config: config,
		logger: l,
		handler: h,
	}
}

func (app *FastHTTPService) Name() string {
	return app.name
}

func (app *FastHTTPService) StartUp() error {
	s := &fasthttp.Server{
		Handler: app.handler,
		Name:    app.name,
	}
	app.listener = app.getListener(app.config.ListenerAddr())

	if err := s.Serve(app.listener); err != nil {
		app.logger.Fatalf("Error when serving incoming connections: %s", err)
		return err
	}
	app.logger.Infof("listening on %s\n", app.config.ListenerAddr())
	return nil
}

func (app *FastHTTPService) getListener(listenAddr string) net.Listener {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ln, err := net.Listen("tcp4", listenAddr);
	if err != nil {
		app.logger.Fatal(err.Error())
	}
	return ln
}

func (app *FastHTTPService) ShutDown() error {
	return app.listener.Close()
}