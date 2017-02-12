package gofury

type HealthCheck interface {
	HealthCheckName() string
	CheckHealth() bool
}

