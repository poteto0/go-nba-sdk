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

func Test_GetIstStandings(t *testing.T) {
	t.Run("can get ist standings w/o params", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.IstStandingsPath + "?LeagueID=00&Season=2025-26&Section=group"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleIstStandingsResponse),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetIstStandings(provider, nil)

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
		assert.Equal(t, 1, len(result.Contents.Teams))
		assert.Equal(t, "Hawks", result.Contents.Teams[0].TeamName)
	})

	t.Run("can get ist standings w/o section", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.IstStandingsPath + "?LeagueID=00&Season=2025-26&Section=group"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleIstStandingsResponse),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetIstStandings(provider, &types.IstStandingsParams{
			LeagueID: "00",
			Season:   "2025-26",
		})

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
		assert.Equal(t, 1, len(result.Contents.Teams))
		assert.Equal(t, "Hawks", result.Contents.Teams[0].TeamName)
	})

	t.Run("network error is w/o status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.IstStandingsPath + "?LeagueID=00&Season=2025-26&Section=group"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewErrorResponder(assert.AnError),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetIstStandings(provider, nil)

		// Assert
		assert.NotNil(t, result.Error)
		assert.Equal(t, 0, result.StatusCode)
	})

	t.Run("response parse error w/ http status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.IstStandingsPath + "?LeagueID=00&Season=2025-26&Section=group"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, "hello"),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetIstStandings(provider, nil)

		// Assert
		assert.NotNil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("if invalid section, return error", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetIstStandings(provider, &types.IstStandingsParams{
			Section: "invalid",
		})

		// Assert
		assert.NotNil(t, result.Error)
	})
}
