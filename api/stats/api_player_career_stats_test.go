package stats_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/api/stats"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_GetPlayerCareerStats(t *testing.T) {
	t.Run("can get player career stats", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayerCareerStatsPath,
			httpmock.NewStringResponder(200, samples.SamplePlayerCareerStats),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetPlayerCareerStats(provider, &types.PlayerCareerStatsParams{
			PlayerID: "203076",
		})

		// Assert
		assert.NoError(t, result.Error)
		assert.NotNil(t, result.Contents)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("network error is w/o status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayerCareerStatsPath,
			httpmock.NewErrorResponder(assert.AnError),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetPlayerCareerStats(provider, &types.PlayerCareerStatsParams{
			PlayerID: "203076",
		})

		// Assert
		assert.Error(t, result.Error)
		assert.Equal(t, 0, result.StatusCode)
	})

	t.Run("response parse error w/ http status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayerCareerStatsPath,
			httpmock.NewStringResponder(200, "invalid"),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetPlayerCareerStats(provider, &types.PlayerCareerStatsParams{
			PlayerID: "203076",
		})

		// Assert
		assert.Error(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("invalid resultSet", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayerCareerStatsPath,
			httpmock.NewStringResponder(200, samples.SampleInvalidPlayerCareerStats),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetPlayerCareerStats(provider, &types.PlayerCareerStatsParams{
			PlayerID: "203076",
		})

		// Assert
		assert.Error(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("player id is required", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetPlayerCareerStats(provider, nil)

		// Assert
		assert.Error(t, result.Error)
	})
}
