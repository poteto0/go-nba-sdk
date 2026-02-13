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

func Test_GetScheduleLeagueV2(t *testing.T) {
	t.Run("can get schedule league v2", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.ScheduleLeagueV2Path + "?LeagueID=00&Season=2024-25"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleScheduleLeagueV2Response),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetScheduleLeagueV2(provider, &types.ScheduleLeagueV2Params{
			LeagueID: "00",
			Season:   "2024-25",
		})

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
		assert.Equal(t, "2024-25", result.Contents.LeagueSchedule.SeasonYear)
		assert.Equal(t, 1, len(result.Contents.LeagueSchedule.GameDates))
		assert.Equal(t, "10/04/2024 00:00:00", result.Contents.LeagueSchedule.GameDates[0].GameDate)
		assert.Equal(t, 1, len(result.Contents.LeagueSchedule.GameDates[0].Games))
		assert.Equal(t, "0012400001", result.Contents.LeagueSchedule.GameDates[0].Games[0].GameID)
	})

	t.Run("can get schedule league v2 with empty params", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.ScheduleLeagueV2Path + "?LeagueID=00&Season=2025-26"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleScheduleLeagueV2Response),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetScheduleLeagueV2(provider, &types.ScheduleLeagueV2Params{})

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("can get schedule league v2 with only LeagueID", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.ScheduleLeagueV2Path + "?LeagueID=01&Season=2025-26"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleScheduleLeagueV2Response),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetScheduleLeagueV2(provider, &types.ScheduleLeagueV2Params{LeagueID: "01"})

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("can get schedule league v2 with only Season", func(t *testing.T) {
		// Arrange
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.ScheduleLeagueV2Path + "?LeagueID=00&Season=2023-24"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, samples.SampleScheduleLeagueV2Response),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetScheduleLeagueV2(provider, &types.ScheduleLeagueV2Params{Season: "2023-24"})

		// Assert
		assert.Nil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})

	t.Run("network error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.ScheduleLeagueV2Path + "?LeagueID=00&Season=2025-26"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewErrorResponder(assert.AnError),
		)

		provider := newProviderForTest()

		result := stats.GetScheduleLeagueV2(provider, nil)

		assert.NotNil(t, result.Error)
	})

	t.Run("error handling", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.ScheduleLeagueV2Path + "?LeagueID=00&Season=2025-26"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(500, "Internal Server Error"),
		)

		provider := newProviderForTest()

		result := stats.GetScheduleLeagueV2(provider, nil)

		assert.NotNil(t, result.Error)
		assert.Equal(t, 500, result.StatusCode)
	})

	t.Run("parse error", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		path := constants.StatsBaseUrl + constants.ScheduleLeagueV2Path + "?LeagueID=00&Season=2025-26"
		httpmock.RegisterResponder(
			"GET",
			path,
			httpmock.NewStringResponder(200, "invalid json"),
		)

		provider := newProviderForTest()

		result := stats.GetScheduleLeagueV2(provider, nil)

		assert.NotNil(t, result.Error)
		assert.Equal(t, 200, result.StatusCode)
	})
}
