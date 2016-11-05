package gofury

import "github.com/kelseyhightower/envconfig"

type Application interface {
	Name() string
	LoadConfig(cfg interface{}) error
	RegisterServices(services ...Service)
	StartUp() error
	ShutDown() error
}

type BaseApplication struct {
	services []Service
}

func (app *BaseApplication) Name() string {
	return "Overwrite Me"
}

func (app *BaseApplication) LoadConfig(cfg interface{}) error {
	return envconfig.Process("", cfg);
}

func (app *BaseApplication) RegisterServices(services ...Service) {
	app.services = append(app.services, services...)
}

func (app *BaseApplication) StartUp() {
	for _, service := range app.services {
		service.StartUp()
	}
}

func (app *BaseApplication) ShutDown() {
	for _, service := range app.services {
		service.ShutDown()
	}
}

