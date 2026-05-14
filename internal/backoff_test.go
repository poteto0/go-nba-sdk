package internal

import (
	"testing"
	"time"

	"github.com/poteto0/go-nba-sdk/types"
)

func TestBackoffManager(t *testing.T) {
	config := types.BackoffConfig{
		Mode:           "exponential",
		MaxRetries:     2,
		InitialDelayMs: 10,
		MaxDelayMs:     50,
	}

	bm := NewBackoffManager(config)

	start := time.Now()
	err := bm.Backoff()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	duration := time.Since(start)
	if duration < 10*time.Millisecond {
		t.Errorf("Expected delay at least 10ms, got %v", duration)
	}

	start = time.Now()
	err = bm.Backoff()
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	duration = time.Since(start)
	if duration < 10*time.Millisecond {
		t.Errorf("Expected delay at least 10ms, got %v", duration)
	}

	err = bm.Backoff()
	if err != types.ErrOverMaxRetries {
		t.Errorf("Expected ErrOverMaxRetries, got %v", err)
	}
}
