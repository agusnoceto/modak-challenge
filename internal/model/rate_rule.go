package model

import "time"

type RateRule struct {
	Name     string
	Interval time.Duration
	Limit    int
}
