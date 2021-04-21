package service

import (
	pb "github.com/meateam/"
)

// Service is the structure used for handling
type Service struct {
	spikeClient spb.SpikeClient
	controller  Controller
	logger      *logrus.Logger
	grantType   string
	audience    string
	approvalURL string
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
