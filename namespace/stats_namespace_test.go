package namespace_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/namespace"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_CreateStatsNamespace(t *testing.T) {
	// Act
	statsNamespace := namespace.NewStatsNamespace(newProviderForTest())

	// Assert
	assert.NotNil(t, statsNamespace)
}

func Test_Stats_GetPlayerCareerStats(t *testing.T) {
	t.Run("can get player career stats", func(t *testing.T) {
		// Arrange
		sn := namespace.NewStatsNamespace(newProviderForTest())

		// Act
		result := sn.GetPlayerCareerStats(&types.PlayerCareerStatsParams{
			PlayerID: "203076",
		})

		// Assert
		assert.NoError(t, result.Error)
		assert.NotNil(t, result.Contents)
		assert.Equal(t, 200, result.StatusCode)
	})
}
