package notification

import (
	"github.com/agusnoceto/modak-challenge/internal/model"
	"github.com/agusnoceto/modak-challenge/internal/notification/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceTestSuite struct {
	suite.Suite
	repository *mocks.Repository
	gateway    Gateway
	limiter    *mocks.Limiter
	service    *RateLimitedService
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

func (s *ServiceTestSuite) SetupTest() {
	s.repository = mocks.NewRepository(s.T())
	s.gateway = *NewGateway()
	s.limiter = mocks.NewLimiter(s.T())
	s.service = NewRateLimitingService(s.repository, s.gateway, s.limiter)
}

func (s *ServiceTestSuite) TestSendAllowed() {
	email := "test@email.com"
	msg := "test message"

	s.limiter.On("Allow", model.MessageKeyNews, email).Return(true, nil).Times(1)
	s.repository.On("Insert", model.MessageKeyNews, email, msg).Return(nil).Times(1)

	err := s.service.Send(model.MessageKeyNews, email, msg)
	assert.NoError(s.T(), err)

}

func (s *ServiceTestSuite) TestSendNotAllowed() {
	email := "test@email.com"
	msg := "test message"

	s.limiter.On("Allow", model.MessageKeyNews, email).Return(false, nil).Times(1)

	err := s.service.Send(model.MessageKeyNews, email, msg)
	assert.Error(s.T(), err)

}
