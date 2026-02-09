package namespace_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/api/live"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/poteto0/go-nba-sdk/namespace"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_CreateLiveNamespace(t *testing.T) {
	// Act
	liveNamespace := namespace.NewLiveNamespace(newProviderForTest())

	// Assert
	assert.NotNil(t, liveNamespace)
}

func Test_Live_GetScoreBoard(t *testing.T) {
	t.Run("can get score board", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.LiveBaseUrl+constants.ScoreBoardPath+"/todaysScoreboard_00.json",
			httpmock.NewStringResponder(200, samples.SampleScoreBoardResponse),
		)

		// Arrange
		sl := namespace.NewLiveNamespace(newProviderForTest())

		// Act
		result := sl.GetScoreBoard(nil)

		// Assert
		assert.NotNil(t, result.Contents)
		assert.NotEmpty(t, result.Contents.Scoreboard.Games)
		assert.Equal(t, 200, result.Contents.Meta.Code)
	})
}

func Test_Live_GetBoxScore(t *testing.T) {
	t.Run("can get box score", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.LiveBaseUrl+constants.BoxScorePath+"/boxscore_0022500733.json",
			httpmock.NewStringResponder(200, samples.SampleScoreBoardResponse),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetBoxScore(provider, &types.BoxScoreParams{
			GameID: "0022500733",
		})

		// Assert
		assert.NotNil(t, result.Contents)
		assert.NoError(t, result.Error)
	})
}
