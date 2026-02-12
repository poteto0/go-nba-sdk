---
sidebar_position: 3
---

# Standings (IST)

## ⚡ Quick Start

```go title="main.go"
func main() {
	client := gns.NewClient(nil)
	result := client.Stats.GetStandings(&types.IstStandingsParams{
		LeagueID: "00",
		Season:   "2023-24",
	})

	if result.Error != nil {
		panic(result.Error)
	}

	for _, record := range result.Contents.Standings {
		fmt.Printf("Team: %s, Conference: %s, Group: %s, Rank: %d\n",
			record.TeamName,
			record.Conference,
			record.Group,
			record.PlayoffRank,
		)
	}
}
```

## 📜 Details

### Arg

```go title="arg.go"
type IstStandingsParams struct {
	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "2023-24"
	Season string `url:"Season"`
}
```

### Response

[`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/response_ist_standings.go)

```
