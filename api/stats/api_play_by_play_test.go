package stats_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/api/stats"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_GetPlayByPlay(t *testing.T) {
	t.Run("can get play by play", func(t *testing.T) {
		// Arrange
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayByPlayV3Path,
			httpmock.NewStringResponder(200, samples.SampleLivePlayByPlayResponse),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetPlayByPlay(provider, &types.PlayByPlayParams{
			GameID: "0022000001",
		})

		// Assert
		assert.NoError(t, result.Error)
		assert.Equal(t, http.StatusOK, result.StatusCode)

		game := result.Contents.Game
		assert.Equal(t, "0022000001", game.GameID)
		assert.NotEmpty(t, game.Actions)
		assert.Equal(t, 1, len(game.Actions))

		action := game.Actions[0]
		assert.Equal(t, 2, action.ActionNumber)
		assert.Equal(t, "period", action.ActionType)
		assert.Equal(t, "Period Start", action.Description)
	})

	t.Run("can get play by play with custom periods", func(t *testing.T) {
		// Arrange
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayByPlayV3Path,
			httpmock.NewStringResponder(200, samples.SampleLivePlayByPlayResponse),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetPlayByPlay(provider, &types.PlayByPlayParams{
			GameID:      "0022000001",
			StartPeriod: 2,
			EndPeriod:   3,
		})

		// Assert
		assert.NoError(t, result.Error)
		assert.Equal(t, http.StatusOK, result.StatusCode)
	})

	t.Run("defaults StartPeriod to 1 and EndPeriod to 4", func(t *testing.T) {
		// Arrange
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayByPlayV3Path,
			httpmock.NewStringResponder(200, samples.SampleLivePlayByPlayResponse),
		)

		provider := newProviderForTest()

		params := &types.PlayByPlayParams{
			GameID: "0022000001",
		}

		// Act
		_ = stats.GetPlayByPlay(provider, params)

		// Assert - verify defaults were applied
		assert.Equal(t, 1, params.StartPeriod)
		assert.Equal(t, 4, params.EndPeriod)
	})

	t.Run("nil params returns error", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetPlayByPlay(provider, nil)

		// Assert
		assert.Error(t, result.Error)
		assert.Equal(t, 0, result.StatusCode)
	})

	t.Run("empty GameID returns error", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := stats.GetPlayByPlay(provider, &types.PlayByPlayParams{})

		// Assert
		assert.Error(t, result.Error)
		assert.Equal(t, 0, result.StatusCode)
	})

	t.Run("network error is w/o status code", func(t *testing.T) {
		// Arrange
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayByPlayV3Path,
			httpmock.NewErrorResponder(assert.AnError),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetPlayByPlay(provider, &types.PlayByPlayParams{
			GameID: "0022000001",
		})

		// Assert
		assert.Error(t, result.Error)
		assert.Equal(t, 0, result.StatusCode)
	})

	t.Run("response parse error w/ http status code", func(t *testing.T) {
		// Arrange
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.PlayByPlayV3Path,
			httpmock.NewStringResponder(200, "invalid json"),
		)

		provider := newProviderForTest()

		// Act
		result := stats.GetPlayByPlay(provider, &types.PlayByPlayParams{
			GameID: "0022000001",
		})

		// Assert
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.Error(t, result.Error)
	})
}
