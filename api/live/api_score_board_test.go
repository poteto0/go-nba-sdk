package live_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/api/live"
	"github.com/stretchr/testify/assert"
)

func Test_GetScoreBoard(t *testing.T) {
	t.Run("can get score board", func(t *testing.T) {
		// Arrange
		provider := newProviderForTest()

		// Act
		result := live.GetScoreBoard(provider, nil)

		// Assert
		assert.NotNil(t, result.Contents)
		assert.NotEmpty(t, result.Contents.Scoreboard.Games)
		assert.Equal(t, 200, result.Contents.Meta.Code)
	})
}
