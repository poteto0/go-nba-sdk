package types

import "time"

type BackoffConfig struct {
	Mode           string
	MaxRetries     int
	InitialDelayMs float64
	MaxDelayMs     float64
}

type GnsConfig struct {
	// default is 10s
	Timeout time.Duration

	// default is empty string
	ProxyUrl string

	Backoff BackoffConfig
}
