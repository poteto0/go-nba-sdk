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

func Test_GetLeagueStandings(t *testing.T) {
	t.Run("can get league standings w/o args", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.LeagueStandingsPath + "?LeagueID=00&Season=2025-26&SeasonType=Regular+Season&SeasonYear="
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleLeagueStandingsResponse),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetLeagueStandings(provider, nil)

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
		assert.Equal(t, 1, len(result.Contents.Standings))
		assert.Equal(t, "Thunder", result.Contents.Standings[0].TeamName)
		assert.Equal(t, 42, result.Contents.Standings[0].Wins)
	})

	t.Run("can get league standings w only season & leagueId", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.LeagueStandingsPath + "?LeagueID=00&Season=2025-26&SeasonType=Regular+Season&SeasonYear="
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleLeagueStandingsResponse),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetLeagueStandings(provider, &types.LeagueStandingsParams{
			LeagueID: "00",
			Season:   "2025-26",
		})

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
		assert.Equal(t, 1, len(result.Contents.Standings))
		assert.Equal(t, "Thunder", result.Contents.Standings[0].TeamName)
		assert.Equal(t, 42, result.Contents.Standings[0].Wins)
	})

	t.Run("parser error w/ http status code", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.LeagueStandingsPath + "?LeagueID=00&Season=2025-26&SeasonType=Regular+Season&SeasonYear="
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleInvalidLeagueStandingsResponse),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetLeagueStandings(provider, nil)

		// Assert
		assert.NotNil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("response parse error w/ http status code", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.LeagueStandingsPath + "?LeagueID=00&Season=2025-26&SeasonType=Regular+Season&SeasonYear="
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, "hello"),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetLeagueStandings(provider, nil)

		// Assert
		assert.NotNil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("internal server error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.LeagueStandingsPath + "?LeagueID=00&Season=2025-26&SeasonType=Regular+Season&SeasonYear="
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(500, "Internal Server Error"),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetLeagueStandings(provider, nil)

		// Assert
		assert.NotNil(t, result.Error)
		assert.Equal(t, 500, result.StatusCode)
	})

	t.Run("invalid SeasonType returns error", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetLeagueStandings(provider, &types.LeagueStandingsParams{
			SeasonType: "Invalid",
		})

		assert.NotNil(t, result.Error)
		assert.Contains(t, result.Error.Error(), "SeasonType is allowed")
	})

	t.Run("can get league standings with SeasonYear", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		// SeasonYear is passed as query param
		path := constants.StatsBaseUrl + constants.LeagueStandingsPath + "?LeagueID=00&Season=2025-26&SeasonType=Regular+Season&SeasonYear=2025"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleLeagueStandingsResponse),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetLeagueStandings(provider, &types.LeagueStandingsParams{
			SeasonYear: "2025",
		})

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})
}
