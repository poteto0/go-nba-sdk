package types_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestPlayerBoxScoreStatistic_ParseMinutesClock(t *testing.T) {
	// Arrange
	playerBoxScoreStatistic := types.PlayerBoxScoreStatistic{
		CommonBoxScoreStatistic: types.CommonBoxScoreStatistic{
			Minutes: "PT07M11.01S",
		},
	}

	// Act & Assert
	assert.Equal(t, "07:11.01", playerBoxScoreStatistic.MinutesClock())
}

func TestPlayerBoxScoreStatistic_CalculateMin(t *testing.T) {
	t.Run("can calculate playtime", func(t *testing.T) {
		// Arrange
		playerBoxScoreStatistic := types.PlayerBoxScoreStatistic{
			CommonBoxScoreStatistic: types.CommonBoxScoreStatistic{
				Minutes: "PT36M30.00S",
			},
		}

		// Act
		min, ok := playerBoxScoreStatistic.Min()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 36.5, min)
	})

	t.Run("failed to calculate on invalid type str", func(t *testing.T) {
		// Arrange
		playerBoxScoreStatistic := types.PlayerBoxScoreStatistic{
			CommonBoxScoreStatistic: types.CommonBoxScoreStatistic{
				Minutes: "invalid",
			},
		}

		// Act
		_, ok := playerBoxScoreStatistic.Min()

		// Assert
		assert.False(t, ok)
	})

	t.Run("failed to calculate on invalid minutes", func(t *testing.T) {
		// Arrange
		playerBoxScoreStatistic := types.PlayerBoxScoreStatistic{
			CommonBoxScoreStatistic: types.CommonBoxScoreStatistic{
				Minutes: "PTaaM30.00S",
			},
		}

		// Act
		_, ok := playerBoxScoreStatistic.Min()

		// Assert
		assert.False(t, ok)
	})

	t.Run("failed to calculate on invalid seconds", func(t *testing.T) {
		// Arrange
		playerBoxScoreStatistic := types.PlayerBoxScoreStatistic{
			CommonBoxScoreStatistic: types.CommonBoxScoreStatistic{
				Minutes: "PT36Maa.00S",
			},
		}

		// Act
		_, ok := playerBoxScoreStatistic.Min()

		// Assert
		assert.False(t, ok)
	})
}
