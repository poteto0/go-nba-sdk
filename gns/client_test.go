package gns_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/gns"
	"github.com/stretchr/testify/assert"
)

func Test_CreateStatsClient(t *testing.T) {
	// Act
	statsClient := gns.NewStatsClient()

	// Assert
	assert.NotNil(t, statsClient)
}
