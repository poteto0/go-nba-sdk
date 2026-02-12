package types_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestGame_IsGameStart(t *testing.T) {
	tests := []struct {
		name string
		game types.Game
		want bool
	}{
		{
			name: "game status 1 (scheduled) is not start",
			game: types.Game{GameStatus: 1},
			want: false,
		},
		{
			name: "game status 2 (live) is start",
			game: types.Game{GameStatus: 2},
			want: true,
		},
		{
			name: "game status 3 (finished) is start",
			game: types.Game{GameStatus: 3},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.game.IsGameStart())
		})
	}
}

func TestGame_IsOverTime(t *testing.T) {
	tests := []struct {
		name string
		game types.Game
		want bool
	}{
		{
			name: "period 4 is not overtime",
			game: types.Game{Period: 4},
			want: false,
		},
		{
			name: "period 5 is overtime",
			game: types.Game{Period: 5},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.game.IsOverTime())
		})
	}
}

func TestGame_OverTimeNum(t *testing.T) {
	tests := []struct {
		name string
		game types.Game
		want int
	}{
		{
			name: "period 4 returns 0",
			game: types.Game{Period: 4},
			want: 0,
		},
		{
			name: "period 5 returns 1",
			game: types.Game{Period: 5},
			want: 1,
		},
		{
			name: "period 6 returns 2",
			game: types.Game{Period: 6},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.game.OverTimeNum())
		})
	}
}
