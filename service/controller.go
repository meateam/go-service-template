package service

import (
	"context"
)

// Controller is an interface for the business logic of the permit.Service which uses a Store.
type Controller interface {
	CreateHelloWorld(ctx context.Context) (Template, error)
	HealthCheck(ctx context.Context) (bool, error)
}
