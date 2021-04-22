package service

import (
	"context"
	"fmt"
	"time"

	pb "github.com/meateam/go-service-template/proto"
	"github.com/sirupsen/logrus"
)

// Service is the structure used for handling
type Service struct {
	controller Controller
	logger     *logrus.Logger
}

// HealthCheck checks the health of the service, and returns a boolean accordingly.
func (s *Service) HealthCheck(mongoClientPingTimeout time.Duration) bool {
	timeoutCtx, cancel := context.WithTimeout(context.TODO(), mongoClientPingTimeout)
	defer cancel()
	healthy, err := s.controller.HealthCheck(timeoutCtx)
	if err != nil {
		s.logger.Errorf("%v", err)
		return false
	}

	return healthy
}

// NewService creates a Service and returns it.
func NewService(controller Controller, logger *logrus.Logger) Service {
	s := Service{controller: controller, logger: logger}
	return s
}

// CreateHelloWorld is the request handler for creating hello world!.
func (s Service) CreateHelloWorld(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	firstName := req.GetFirstName()
	lastName := req.GetLastName()

	if firstName == "" {
		return nil, fmt.Errorf("firstName is required")
	}

	if lastName == "" {
		return nil, fmt.Errorf("lastName is required")
	}

	return &pb.HelloWorldResponse{}, nil
}
