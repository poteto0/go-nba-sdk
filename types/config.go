package types

import "time"

type GnsConfig struct {
	// default is 10s
	Timeout time.Duration

	// default is empty string
	ProxyUrl string
}
