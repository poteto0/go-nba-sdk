---
sidebar_position: 3
---

# League Standings

get League Standings

## ⚡ Quick Start

```go title="main.go"
func main() {
	client := gns.NewClient(nil)
	result := client.Stats.GetLeagueStandings(&types.LeagueStandingsParams{
		LeagueID: "00",
		Season:   "2025-26",
		SeasonType: "Regular Season",
	})

	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Printf("Season: %s\n", result.Contents.SeasonYear)
	for _, team := range result.Contents.Standings {
		fmt.Printf("Team: %s %s, Record: %s, Rank: %d\n",
			team.TeamCity,
			team.TeamName,
			team.Record,
			team.PlayoffRank,
		)
	}
}
```

output

Season: 2025-26
Team: Oklahoma City Thunder, Record: 42-13, Rank: 1

## 📜 Details

### Arg

```go title="arg.go"
type LeagueStandingsParams struct {
	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "2025-26"
	Season string `url:"Season"`

	// optional default "Regular Season"
	// "Regular Season" or "Pre Season"
	SeasonType string `url:"SeasonType"`

	// optional
	SeasonYear string `url:"SeasonYear"`
}
```

### Response

[`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/response_league_standings.go)

- Nullable fields return `optional.Option[string]`.

  ```go
  result := stats.GetLeagueStandings(client, types.LeagueStandingsParams{})

  for _, team := range result.Contents.Standings {
  	// Example of handling nullable field
  	if v, err := team.ClinchedPlayoffBirth.Take(); err == nil {
  		fmt.Printf("Clinched: %s\n", v)
  	}
  }
  ```