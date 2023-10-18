package notification

import (
	"fmt"
	"github.com/agusnoceto/modak-challenge/internal/model"
	"time"
)

type RateLimiter struct {
	rules      map[model.MessageKey]model.RateRule
	repository Repository
}

func NewRateLimiter(rules map[model.MessageKey]model.RateRule, repository Repository) *RateLimiter {
	return &RateLimiter{
		rules:      rules,
		repository: repository,
	}
}

func (l *RateLimiter) Allow(key model.MessageKey, email string) (bool, error) {
	rule, ok := l.rules[key]
	if !ok {
		return false, fmt.Errorf("missing rule for key: %s", key)
	}

	since := time.Now().Add(-1 * rule.Interval)
	sentMsgs, err := l.repository.GetByEmailSince(key, email, since)
	if err != nil {
		return false, err
	}
	if len(sentMsgs) < rule.Limit {
		return true, nil
	}
	return false, nil
}
