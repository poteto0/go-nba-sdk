package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Player(t *testing.T) {
	t.Run("can check if player is starter", func(t *testing.T) {
		// Arrange
		tests := []struct {
			name   string
			player Player
			want   bool
		}{
			{
				name: "player is starter",
				player: Player{
					Starter: "1",
				},
				want: true,
			},
			{
				name: "player is not starter",
				player: Player{
					Starter: "0",
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Act
				result := tt.player.IsStarter()

				// Assert
				assert.Equal(t, result, tt.want)
			})
		}
	})

	t.Run("check on-court status", func(t *testing.T) {
		// Arrange
		tests := []struct {
			name   string
			player Player
			want   bool
		}{
			{
				name: "player is on court",
				player: Player{
					OnCourt: "1",
				},
				want: true,
			},
			{
				name: "player is not on court",
				player: Player{
					OnCourt: "0",
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Act
				result := tt.player.IsOnCourt()

				// Assert
				assert.Equal(t, result, tt.want)
			})
		}
	})

	t.Run("check is played", func(t *testing.T) {
		// Arrange
		tests := []struct {
			name   string
			player Player
			want   bool
		}{
			{
				name: "player is on court",
				player: Player{
					Played: "1",
				},
				want: true,
			},
			{
				name: "player is not on court",
				player: Player{
					Played: "0",
				},
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Act
				result := tt.player.IsPlayed()

				// Assert
				assert.Equal(t, result, tt.want)
			})
		}
	})
}
