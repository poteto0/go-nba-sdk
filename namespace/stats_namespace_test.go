package namespace_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
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
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayerCareerStatsPath,
			httpmock.NewStringResponder(200, samples.SamplePlayerCareerStats),
		)

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

func Test_Stats_GetIstStandings(t *testing.T) {
	t.Run("can get ist standings", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.IstStandingsPath,
			httpmock.NewStringResponder(200, samples.SampleIstStandingsResponse),
		)

		// Arrange
		sn := namespace.NewStatsNamespace(newProviderForTest())

		// Act
		result := sn.GetIstStandings(&types.IstStandingsParams{
			LeagueID: "00",
			Season:   "2025-26",
		})

		// Assert
		assert.NoError(t, result.Error)
		assert.NotNil(t, result.Contents)
		assert.Equal(t, 200, result.StatusCode)
	})
}
