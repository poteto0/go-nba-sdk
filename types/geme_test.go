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
