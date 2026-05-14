package internal

import (
	"crypto/rand"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/poteto0/go-nba-sdk/types"
)

type BackoffManager struct {
	config    types.BackoffConfig
	retries   int
	lastDelay float64
	lock      sync.Mutex
}

func NewBackoffManager(config types.BackoffConfig) *BackoffManager {
	return &BackoffManager{
		config:    config,
		retries:   0,
		lastDelay: 0,
	}
}

func (b *BackoffManager) Backoff() error {
	b.lock.Lock()
	defer b.lock.Unlock()

	if b.retries >= b.config.MaxRetries {
		return types.ErrOverMaxRetries
	}

	currentDelay := b.calculateBackOffDelay()

	currentDelay = math.Max(currentDelay, b.config.InitialDelayMs)
	currentDelay = math.Min(currentDelay, b.config.MaxDelayMs)

	b.retries++
	b.lastDelay = currentDelay

	time.Sleep(time.Duration(currentDelay) * time.Millisecond)

	return nil
}

func (b *BackoffManager) calculateBackOffDelay() float64 {
	switch b.config.Mode {
	case "exponential":
		return b.calculateExponentialBackOffDelay()
	default:
		return b.config.InitialDelayMs
	}
}

func (b *BackoffManager) calculateExponentialBackOffDelay() float64 {
	if b.retries == 0 {
		return float64(0)
	}

	return math.Min(b.lastDelay+(randomF64(1.0)-0.5)*b.lastDelay, b.config.MaxDelayMs)
}

func randomF64(max float64) float64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return float64(nBig.Int64() / int64(max))
}
