package gofury

type Service interface {
	Name()
	StartUp()
	ShutDown()
}