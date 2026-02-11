package types_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestGame_JudgeTheGameIsFinished(t *testing.T) {
	// Arrange
	tests := []struct {
		name string
		game types.Game
		want bool
	}{
		{
			name: "game is not finished",
			game: types.Game{
				GameStatus: 1,
			},
			want: false,
		},
		{
			name: "game is finished",
			game: types.Game{
				GameStatus: 3,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act
			result := tt.game.IsFinished()

			// Assert
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestGame_ParseGameClock(t *testing.T) {
	// Arrange
	game := types.Game{
		GameClock: "PT07M11.01S",
	}

	// Act & Assert
	assert.Equal(t, "07:11.01", game.Clock())
}

func TestGame_ParseClockMinutes(t *testing.T) {
	t.Run("parse valid minutes", func(t *testing.T) {
		// Arrange
		game := types.Game{
			GameClock: "PT07M11.01S",
		}

		// Act
		minutes, ok := game.ClockMinutes()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 7.0, minutes)
	})

	t.Run("not ok on invalid minutes", func(t *testing.T) {
		// Arrange
		game := types.Game{
			GameClock: "PTaaM11.01S",
		}

		// Act
		_, ok := game.ClockMinutes()

		// Assert
		assert.False(t, ok)
	})
}

func TestGame_ParseClockSeconds(t *testing.T) {
	t.Run("parse valid seconds", func(t *testing.T) {
		// Arrange
		game := types.Game{
			GameClock: "PT07M11.01S",
		}

		// Act
		seconds, ok := game.ClockSeconds()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 11.01, seconds)
	})

	t.Run("not ok on invalid seconds", func(t *testing.T) {
		// Arrange
		game := types.Game{
			GameClock: "PT07Maa.01S",
		}

		// Act
		_, ok := game.ClockSeconds()

		// Assert
		assert.False(t, ok)
	})
}
