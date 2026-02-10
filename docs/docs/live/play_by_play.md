---
sidebar_position: 4
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
	result := client.Live.GetPlayByPlay(
		&types.PlayByPlayParams{
			GameID: "0022000001",
		},
	)

	actions := result.Contents.Game.Actions
	for i, action := range actions {
		if i >= 5 {
			break
		}

		fmt.Printf(
			"%dQ (%s) %s: %s\n",
			action.Period,
			action.Clock,
			action.PlayerNameI,
			action.ActionType,
		)
	}
}
```

get play by play

```
1Q (PT12M00.00S) : period
1Q (PT11M58.00S) : jumpball
1Q (PT11M50.00S) D. Jordan: turnover
1Q (PT11M38.00S) K. Irving: foul
1Q (PT11M38.00S) S. Curry: freethrow
```

## 📜 Details

### Arg

```go title="arg.go"
type PlayByPlayParams struct {
	// required
	GameID string
}
```

### Response

[`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/response_live_play_by_play.go)
