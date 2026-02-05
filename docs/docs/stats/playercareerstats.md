---
sidebar_position: 2
---

# PlayerCareerStats

## ⚡ Quick Start

```go title="main.go"
package main

import (
	"fmt"

	"github.com/poteto0/go-nba-sdk/api/stats"
	"github.com/poteto0/go-nba-sdk/gns"
)

func main() {
	client := gns.NewClient()
	result := stats.GetPlayerCareerStats(client, stats.PlayerCareerStatsParams{
		PlayerID: "203076",
	})

	fmt.Println("======================= RS ==========================")
	for _, content := range result.Contents.SeasonTotalsRegularSeason {
		pts, _ := content.Pts.Take()
		fmt.Printf("Season: %s, PPG: %.2f\n",
			content.SeasonID,
			pts,
		)
	}

	fmt.Println("======================= PO ==========================")
	for _, content := range result.Contents.SeasonTotalsPostSeason {
		pts, _ := content.Pts.Take()
		fmt.Printf("Season: %s, PPG: %.2f\n",
			content.SeasonID,
			pts,
		)
	}

	fmt.Println("======================= COLLEGE ==========================")
	for _, content := range result.Contents.SeasonTotalsCollegeSeason {
		pts, _ := content.Pts.Take()
		fmt.Printf("Season: %s, PPG: %.2f\n",
			content.SeasonID,
			pts,
		)
	}

	fmt.Println("======================= RANK ==============================")
	for _, content := range result.Contents.SeasonRankingsRegularSeason {

		if content.RankPgPts.IsNone() {
			fmt.Printf("Season: %s, PPG Rank: -\n",
				content.SeasonID,
			)
			continue
		}

		v, _ := content.RankPgPts.Take()
		fmt.Printf("Season: %s, PPG Rank: %d\n",
			content.SeasonID,
			v,
		)
	}
}
```

## 📜 Details

- [`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/player_career_stats.go)

- Many columns return gooptional.

  ```go
  result := stats.GetPlayerCareerStats(client, stats.PlayerCareerStatsParams{
  	PlayerID: "203076",
  })

  for _, content := range result.Contents.SeasonTotalsRegularSeason {
  	// check if pts is present or not
  	pts, err := content.Pts.Take()
  }
  ```
