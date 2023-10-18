package notification

import (
	"fmt"
	"github.com/agusnoceto/modak-challenge/internal/model"
	"github.com/agusnoceto/modak-challenge/internal/notification/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type RateLimiterTestSuite struct {
	suite.Suite
	repository  *mocks.Repository
	rateLimiter *RateLimiter
}

func TestRateLimiterTestSuite(t *testing.T) {
	suite.Run(t, new(RateLimiterTestSuite))
}

func (s *RateLimiterTestSuite) SetupTest() {
	s.repository = mocks.NewRepository(s.T())

	rules := map[model.MessageKey]model.RateRule{}
	rules[model.MessageKeyStatus] = model.RateRule{Interval: time.Minute, Limit: 2}
	rules[model.MessageKeyNews] = model.RateRule{Interval: 24 * time.Hour, Limit: 1}
	rules[model.MessageKeyMarketing] = model.RateRule{Interval: time.Hour, Limit: 3}

	s.rateLimiter = NewRateLimiter(rules, s.repository)
}

func (s *RateLimiterTestSuite) TestAllow() {
	email := "test@email.com"

	s.repository.On("GetByEmailSince", model.MessageKeyStatus, email, mock.AnythingOfType("time.Time")).Return([]model.Message{}, nil)

	allow, err := s.rateLimiter.Allow(model.MessageKeyStatus, email)

	assert.NoError(s.T(), err)
	assert.True(s.T(), allow)

	//generate a list with more results than the limit for this key
	newsRule := s.rateLimiter.rules[model.MessageKeyNews]
	msgs := []model.Message{}
	for i := 0; i <= newsRule.Limit; i++ {
		msgs = append(msgs, model.Message{
			Id:     uuid.New(),
			Email:  email,
			Msg:    fmt.Sprintf("Message %d", i),
			SentAt: time.Now(),
		})
	}

	s.repository.On("GetByEmailSince", model.MessageKeyNews, email, mock.AnythingOfType("time.Time")).Return(msgs, nil)
	allow, err = s.rateLimiter.Allow(model.MessageKeyNews, email)

	assert.NoError(s.T(), err)
	assert.False(s.T(), allow)
}
