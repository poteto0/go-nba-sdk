package main

import (
	"fmt"

	"github.com/poteto0/go-nba-sdk/api/stats"
	"github.com/poteto0/go-nba-sdk/gns"
)

func main() {
	client := gns.NewStatsClient()
	result := stats.GetPlayerCareerStats(client, stats.PlayerCareerStatsParams{
		PlayerID: "203076",
	})

	fmt.Println("======================= RS ==========================")
	for _, content := range result.Contents.SeasonTotalsRegularSeason {
		fmt.Printf("Season: %s, PPG: %.2f\n",
			content.SeasonID,
			content.PTS,
		)
	}

	fmt.Println("======================= PO ==========================")
	for _, content := range result.Contents.SeasonTotalsPostSeason {
		fmt.Printf("Season: %s, PPG: %.2f\n",
			content.SeasonID,
			content.PTS,
		)
	}

	fmt.Println("======================= COLLEGE ==========================")
	for _, content := range result.Contents.SeasonTotalsCollegeSeason {
		fmt.Printf("Season: %s, PPG: %.2f\n",
			content.SeasonID,
			content.PTS,
		)
	}

	fmt.Println("======================= RANK ==============================")
	for _, content := range result.Contents.SeasonRankingsRegularSeason {

		if content.RankPgPts == 0 {
			fmt.Printf("Season: %s, PPG Rank: -\n",
				content.SeasonID,
			)
			continue
		}

		fmt.Printf("Season: %s, PPG Rank: %d\n",
			content.SeasonID,
			content.RankPgPts,
		)
	}
}
