package main

import (
	"fmt"

	"github.com/poteto0/go-nba-sdk/gns"
	"github.com/poteto0/go-nba-sdk/types"
)

func main() {
	client := gns.NewClient()
	result := client.Stats.GetPlayerCareerStats(types.PlayerCareerStatsParams{
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
