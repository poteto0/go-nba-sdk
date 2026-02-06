---
sidebar_position: 2
---

# ScoreBoard

## ⚡ Quick Start

```go title="main.go"
package main

import (
	"fmt"

	"github.com/poteto0/go-nba-sdk/gns"
)

func main() {
	client := gns.NewClient()
	result := client.Live.GetScoreBoard(nil)

	games := result.Contents.Scoreboard.Games
	for _, game := range games {
		fmt.Println("==========")
		if game.IsFinished() {
			fmt.Println("  Final  ")
		} else {
			fmt.Printf("%dQ (%s)\n", game.Period, game.GameClock)
		}
		fmt.Printf("%s | %s\n", game.HomeTeam.TeamTricode, game.AwayTeam.TeamTricode)
		fmt.Printf("%d | %d\n", game.HomeTeam.Score, game.AwayTeam.Score)
	}
	fmt.Println("==========")
}
```

get scoreboard

```
==========
  Final
DET | WAS
117 | 126
==========
  Final
ORL | BKN
118 | 98
==========
  Final
ATL | UTA
121 | 119
==========
  Final
TOR | CHI
123 | 107
==========
  Final
HOU | CHA
99 | 109
==========
  Final
DAL | SAS
123 | 135
==========
  Final
PHX | GSW
97 | 101
==========
  Final
LAL | PHI
119 | 115
==========
```

## 📜 Details

### Arg

```go title="arg.go"
type ScoreBoardParams struct {
	// optional default "00"
	LeagueID string
}
```

### Response

[`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/live_score_board.go)
