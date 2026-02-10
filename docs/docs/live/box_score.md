---
sidebar_position: 3
---

# BoxScore

## :zap: Quick Start

```go title="main.go"
package main

import (
	"fmt"

	"github.com/poteto0/go-nba-sdk/gns"
	"github.com/poteto0/go-nba-sdk/types"
)

func main() {
	client := gns.NewClient(nil)
	result := client.Live.GetBoxScore(
		&types.BoxScoreParams{
			GameID: "0022500733",
		},
	)

	game := result.Contents.Game
	homeTeam := game.HomeTeam

	fmt.Printf("=============== %s ==================\n", homeTeam.TeamTricode)

	players, _ := homeTeam.Players.Take()
	for _, player := range players {
		stats, _ := player.Statistics.Take()
		pts, _ := stats.Pts.Take()
		ast, _ := stats.Ast.Take()
		reb, _ := stats.Reb.Take()

		fmt.Printf(
			"%s: %d pts/ %d ast/ %d reb\n",
			player.NameI,
			pts,
			ast,
			reb,
		)
	}
}
```

get box score

```
=============== DET ==================
A. Thompson: 13 pts/ 3 ast/ 9 reb
I. Stewart: 4 pts/ 0 ast/ 2 reb
J. Duren: 4 pts/ 0 ast/ 3 reb
D. Robinson: 21 pts/ 2 ast/ 1 reb
C. Cunningham: 30 pts/ 8 ast/ 8 reb
R. Holland II: 11 pts/ 2 ast/ 10 reb
C. LeVert: 7 pts/ 3 ast/ 2 reb
M. Sasser: 8 pts/ 3 ast/ 2 reb
J. Green: 8 pts/ 0 ast/ 3 reb
K. Huerter: 0 pts/ 1 ast/ 1 reb
P. Reed: 7 pts/ 0 ast/ 5 reb
W. Moore Jr.: 4 pts/ 0 ast/ 1 reb
B. Klintman: 0 pts/ 0 ast/ 0 reb
C. Lanier: 0 pts/ 0 ast/ 0 reb
T. Harris: 0 pts/ 0 ast/ 0 reb
D. Jenkins: 0 pts/ 0 ast/ 0 reb
T. Smith: 0 pts/ 0 ast/ 0 reb
```

## 📜 Details

### Arg

```go title="arg.go"
type BoxScoreParams struct {
	// required
	GameID string
}
```

### Response

[`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/response_live_box_score.go)
