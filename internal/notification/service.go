package notification

import (
	"fmt"
	"github.com/agusnoceto/modak-challenge/internal/model"
)

type gateway struct{}

func NewGateway() *gateway {
	return &gateway{}
}
func (g *gateway) Send(email string, msg string) {
	fmt.Printf("Sending msg to user %s: %s\n", email, msg)
}

type rateLimitedService struct {
	repository Repository
	gateway    gateway
	limiter    Limiter
}

func NewRateLimitingService(repository Repository, gateway gateway, limiter Limiter) *rateLimitedService {
	return &rateLimitedService{
		repository: repository,
		gateway:    gateway,
		limiter:    limiter,
	}
}

func (s *rateLimitedService) Send(key model.MessageKey, email string, msg string) error {
	allow, err := s.limiter.Allow(key, email)
	if err != nil {
		return err
	}
	if !allow {
		return fmt.Errorf("rate limit exceeded for key %s and user %s\n", key, email)
	}

	s.gateway.Send(email, msg)
	return s.repository.Insert(key, email, msg)
}
