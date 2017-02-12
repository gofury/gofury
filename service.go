package gofury

type Service interface {
	Name() string
	StartUp() error
	ShutDown() error
}