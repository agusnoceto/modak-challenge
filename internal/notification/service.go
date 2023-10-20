package notification

import (
	"fmt"
	"github.com/agusnoceto/modak-challenge/internal/model"
)

type Gateway struct{}

func NewGateway() *Gateway {
	return &Gateway{}
}
func (g *Gateway) Send(email string, msg string) {
	fmt.Printf("Sending msg to user %s: %s\n", email, msg)
}

type RateLimitedService struct {
	repository Repository
	gateway    Gateway
	limiter    Limiter
}

func NewRateLimitingService(repository Repository, gateway Gateway, limiter Limiter) *RateLimitedService {
	return &RateLimitedService{
		repository: repository,
		gateway:    gateway,
		limiter:    limiter,
	}
}

func (s *RateLimitedService) Send(key model.MessageKey, email string, msg string) error {
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
