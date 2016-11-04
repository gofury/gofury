package gofury

const (
	// HealthCheckerHealthy means the service is healthy
	HealthCheckerHealthy = "Healthy"
	// HealthCheckerUnhealthy means the service is unhealthy
	HealthCheckerUnhealthy = "Unhealthy"
)

// HealthChecker interface.
type HealthChecker interface {
	CheckHealth() string
	HealthCheckerName() string
}
