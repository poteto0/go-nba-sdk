# go-nba-sdk

NBA API SDK for golang

[:memo: document](https://docs-go-nba-sdk.poteto-mahiro.com/)

## :zap: Quick Start

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
