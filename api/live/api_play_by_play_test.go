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

func Test_GetPlayByPlay(t *testing.T) {
	t.Run("can get play by play", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.LiveBaseUrl+constants.PlayByPlayPath+"/playbyplay_0022000001.json",
			httpmock.NewStringResponder(200, samples.SampleLivePlayByPlayResponse),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetPlayByPlay(provider, &types.PlayByPlayParams{
			GameID: "0022000001",
		})

		// Assert
		assert.NotNil(t, result.Contents)
		assert.Equal(t, http.StatusOK, result.StatusCode)
		assert.NoError(t, result.Error)

		// assert contents
		game := result.Contents.Game
		assert.Equal(t, "0022000001", game.GameID)
		assert.NotEmpty(t, game.Actions)
		assert.Equal(t, 1, len(game.Actions))

		action := game.Actions[0]
		assert.Equal(t, 2, action.ActionNumber)
		assert.Equal(t, "period", action.ActionType)
		assert.Equal(t, "Period Start", action.Description)
	})

	t.Run("network error is w/o status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.LiveBaseUrl+constants.PlayByPlayPath+"/playbyplay_0022000001.json",
			httpmock.NewErrorResponder(assert.AnError),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetPlayByPlay(provider, &types.PlayByPlayParams{
			GameID: "0022000001",
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
			constants.LiveBaseUrl+constants.PlayByPlayPath+"/playbyplay_0022000001.json",
			httpmock.NewStringResponder(200, "invalid"),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetPlayByPlay(provider, &types.PlayByPlayParams{
			GameID: "0022000001",
		})

		// Assert
		assert.Equal(t, 200, result.StatusCode)
		assert.Error(t, result.Error)
	})

	t.Run("nil arg returns error", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetPlayByPlay(provider, nil)

		// Assert
		assert.Error(t, result.Error)
	})

	t.Run("gameID is required", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetPlayByPlay(provider, &types.PlayByPlayParams{})

		// Assert
		assert.Error(t, result.Error)
	})
}
