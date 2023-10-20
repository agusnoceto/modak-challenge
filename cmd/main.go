package main

import (
	"github.com/agusnoceto/modak-challenge/internal/model"
	"github.com/agusnoceto/modak-challenge/internal/notification"
	"time"
)

func main() {
	gateway := notification.NewGateway()
	repository := notification.NewVolatileRepository()

	rules := map[model.MessageKey]model.RateRule{}
	rules[model.MessageKeyStatus] = model.RateRule{Interval: time.Minute, Limit: 2}
	rules[model.MessageKeyNews] = model.RateRule{Interval: 24 * time.Hour, Limit: 1}
	rules[model.MessageKeyMarketing] = model.RateRule{Interval: time.Hour, Limit: 3}

	limiter := notification.NewRateLimiter(rules, repository)
	service := notification.NewRateLimitingService(repository, *gateway, limiter)

	service.Send(model.MessageKeyNews, "test@email.com", "Message 1")
	service.Send(model.MessageKeyNews, "test@email.com", "Message 2")

}
