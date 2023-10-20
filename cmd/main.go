package main

import (
	"fmt"
	"github.com/agusnoceto/modak-challenge/internal/model"
	"github.com/agusnoceto/modak-challenge/internal/notification"
	"github.com/agusnoceto/modak-challenge/ui"
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

	ui.PrintWelcomeMessage()

	keepSending := true

	for keepSending {
		ui.PrintDelimiter()

		key, email, msg := ui.ReadValues()
		err := service.Send(key, email, msg)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err))
		}
		keepSending = ui.SendAnotherMessage()
	}
	ui.PrintGoodBye()
}
