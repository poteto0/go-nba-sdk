--- 
sidebar_position: 5
---

# Schedule League V2

get schedules of the league

## ⚡ Quick Start

```go title="main.go"
func main() {
	client := gns.NewClient(nil)
	result := client.Stats.GetScheduleLeagueV2(&types.ScheduleLeagueV2Params{
		LeagueID: "00",
		Season:   "2024-25",
	})

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Printf("Season: %s\n", result.Contents.LeagueSchedule.SeasonYear)
	for _, gameDate := range result.Contents.LeagueSchedule.GameDates {
		fmt.Printf("Date: %s\n", gameDate.GameDate)
		for _, game := range gameDate.Games {
			fmt.Printf("  Game: %s vs %s, Status: %s\n",
				game.HomeTeam.TeamName,
				game.AwayTeam.TeamName,
				game.GameStatusText,
			)
		}
	}
}
```

## 📜 Details

### Arg

```go title="arg.go"
type ScheduleLeagueV2Params struct {
	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "2025-26"
	Season string `url:"Season"`
}
```

### Response

[`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/response_schedule_league_v2.go)

```
