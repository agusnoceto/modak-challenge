package main

import (
	"encoding/json"
	"fmt"
	"github.com/agusnoceto/modak-challenge/internal/model"
	"github.com/agusnoceto/modak-challenge/internal/notification"
	"github.com/agusnoceto/modak-challenge/ui"
	"os"
	"time"
)

func main() {
	gateway := notification.NewGateway()
	repository := notification.NewVolatileRepository()

	rules, err := parseRulesJSONFile("rules.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	limiter := notification.NewRateLimiter(rules, repository)
	service := notification.NewRateLimitingService(repository, *gateway, limiter)

	ui.PrintWelcomeMessage()

	keepSending := true

	for keepSending {
		ui.PrintDelimiter()

		key, email, msg := ui.ReadValues(rules)
		err := service.Send(key, email, msg)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err))
		}
		keepSending = ui.SendAnotherMessage()
	}
	ui.PrintGoodBye()
}

func parseRulesJSONFile(filename string) (map[model.MessageKey]model.RateRule, error) {
	type tempRule struct {
		Name     string
		Interval string
		Limit    int
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var rules []tempRule

	if err := json.Unmarshal(content, &rules); err != nil {
		return nil, err
	}

	ruleMap := make(map[model.MessageKey]model.RateRule)
	for _, rule := range rules {
		interval, err := time.ParseDuration(rule.Interval)
		if err != nil {
			return nil, err
		}

		ruleMap[model.MessageKey(rule.Name)] = model.RateRule{
			Name:     rule.Name,
			Interval: interval,
			Limit:    rule.Limit,
		}
	}

	return ruleMap, nil
}
