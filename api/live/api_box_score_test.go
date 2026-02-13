package live_test

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/api/live"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_GetBoxScore(t *testing.T) {
	t.Run("can get box score", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.LiveBaseUrl+constants.BoxScorePath+"/boxscore_0022500733.json",
			httpmock.NewStringResponder(200, samples.SampleLiveBoxScoreResponse),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetBoxScore(provider, &types.BoxScoreParams{
			GameID: "0022500733",
		})

		// Assert
		assert.NotNil(t, result.Contents)
		assert.Equal(t, result.StatusCode, http.StatusOK)
		assert.NoError(t, result.Error)

		// assert contents
		game := result.Contents.Game
		assert.Equal(t, "0022500733", game.GameId)

		assert.NotNil(t, game.Arena)
		assert.Equal(t, "Little Caesars Arena", game.Arena.ArenaName)

		assert.NotNil(t, game.Officials)
		assert.Equal(t, 1, len(*game.Officials))

		homeTeam := game.HomeTeam
		assert.Equal(t, homeTeam.TeamId, 1610612765)

		periods := homeTeam.Periods
		assert.Equal(t, 4, len(periods))

		assert.NotNil(t, homeTeam.Players)
		assert.Equal(t, 1, len(*homeTeam.Players))

		player := (*homeTeam.Players)[0]
		assert.Equal(t, "Ausar Thompson", player.Name)

		assert.NotNil(t, player.Statistics)
		assert.NotNil(t, player.Statistics.Pts)
		assert.Equal(t, 13, *player.Statistics.Pts)
		assert.NotNil(t, player.Statistics.Plus)
		assert.Equal(t, 79.0, *player.Statistics.Plus)

		assert.NotNil(t, homeTeam.Statistics)
		assert.NotNil(t, homeTeam.Statistics.Pts)
		assert.Equal(t, 117, *homeTeam.Statistics.Pts)
	})

	t.Run("network error is w/o status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.LiveBaseUrl+constants.BoxScorePath+"/boxscore_0022500733.json",
			httpmock.NewErrorResponder(assert.AnError),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetBoxScore(provider, &types.BoxScoreParams{
			GameID: "0022500733",
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
			constants.LiveBaseUrl+constants.BoxScorePath+"/boxscore_0022500733.json",
			httpmock.NewStringResponder(200, "hello"),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetBoxScore(provider, &types.BoxScoreParams{
			GameID: "0022500733",
		})

		// Assert
		assert.Equal(t, 200, result.StatusCode)
		assert.Error(t, result.Error)
	})

	t.Run("nil arg returns error", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetBoxScore(provider, nil)

		// Assert
		assert.Error(t, result.Error)
	})

	t.Run("gameID is required", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetBoxScore(provider, &types.BoxScoreParams{})

		// Assert
		assert.Error(t, result.Error)
	})
}
