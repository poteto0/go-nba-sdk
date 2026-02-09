package live_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/api/live"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/stretchr/testify/assert"
)

func Test_GetScoreBoard(t *testing.T) {
	t.Run("can get score board", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.LiveBaseUrl+constants.ScoreBoardPath+"/todaysScoreboard_00.json",
			httpmock.NewStringResponder(200, samples.SampleScoreBoardResponse),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetScoreBoard(provider, nil)

		// Assert
		assert.NotNil(t, result.Contents)
		assert.NotEmpty(t, result.Contents.Scoreboard.Games)
		assert.Equal(t, 200, result.Contents.Meta.Code)
	})

	t.Run("response parse error w/ http status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.LiveBaseUrl+constants.ScoreBoardPath+"/todaysScoreboard_00.json",
			httpmock.NewStringResponder(200, "hello"),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetScoreBoard(provider, nil)

		// Assert
		assert.Equal(t, 200, result.StatusCode)
		assert.Error(t, result.Error)
	})
}
