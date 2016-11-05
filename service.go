package gofury

type Service interface {
	Name() string
	StartUp() error
	ShutDown() error
}

type QueueService interface {
	Service
	SendChan() chan<- *Model
	ReceiveChan() <-chan *Model
	DeleteChan() chan<- string
}