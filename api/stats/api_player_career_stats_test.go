package stats_test

import (
	"fmt"
	"testing"

	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/api/stats"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func newProviderForTest() api.IProvider {
	return api.NewProvider()
}

func Test_GetPlayerCareerStats(t *testing.T) {
	t.Run("can get player career stats", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetPlayerCareerStats(provider, types.PlayerCareerStatsParams{
			PlayerID: "203076",
		})

		// Assert
		fmt.Println(result)
		assert.NoError(t, result.Error)
		assert.NotNil(t, result.Contents)
		assert.Equal(t, 200, result.StatusCode)
	})
}
