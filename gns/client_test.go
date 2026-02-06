package gns_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/gns"
	"github.com/stretchr/testify/assert"
)

func Test_CreateStatsClient(t *testing.T) {
	// Act
	cl := gns.NewClient()

	// Assert
	assert.NotNil(t, cl)
	assert.NotNil(t, cl.Stats)
	assert.NotNil(t, cl.Live)
}
