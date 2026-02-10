package gns_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/gns"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_CreateClient(t *testing.T) {
	t.Run("can create client w/o config", func(t *testing.T) {
		// Act
		cl := gns.NewClient(nil)

		// Assert
		assert.NotNil(t, cl)
		assert.NotNil(t, cl.Stats)
		assert.NotNil(t, cl.Live)
	})

	t.Run("can create client w/o timeout", func(t *testing.T) {
		// Act
		cl := gns.NewClient(&types.GnsConfig{
			Timeout: 0,
		})

		// Assert
		assert.NotNil(t, cl)
		assert.NotNil(t, cl.Stats)
		assert.NotNil(t, cl.Live)
	})
}
