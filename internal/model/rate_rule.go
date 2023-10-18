package model

import "time"

type RateRule struct {
	Interval time.Duration
	Limit    int
}
