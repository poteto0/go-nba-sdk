---
sidebar_position: 4
---

# Standings (IST)

get IST(In-Season-Tournament) Standings

## ⚡ Quick Start

```go title="main.go"
func main() {
	client := gns.NewClient(nil)
	result := client.Stats.GetIstStandings(&types.IstStandingsParams{
		LeagueID: "00",
		Season:   "2025-26",
		Section:  "group", // or "wildcard"
	})

	if result.Error != nil {
		panic(result.Error)
	}

	for _, team := range result.Contents.Teams {
		fmt.Printf("Team: %s, Conference: %s, Group: %s, Rank: %d\n",
			team.TeamName,
			team.Conference,
			team.IstGroup,
			team.IstGroupRank,
		)
	}
}
```

output

Team: Hawks, Conference: East, Group: East Group A, Rank: 2
Team: Celtics, Conference: East, Group: East Group B, Rank: 2
Team: Cavaliers, Conference: East, Group: East Group A, Rank: 3
Team: Pelicans, Conference: West, Group: West Group B, Rank: 5
Team: Bulls, Conference: East, Group: East Group C, Rank: 5
Team: Mavericks, Conference: West, Group: West Group B, Rank: 4
Team: Nuggets, Conference: West, Group: West Group C, Rank: 2
Team: Warriors, Conference: West, Group: West Group C, Rank: 5
Team: Rockets, Conference: West, Group: West Group C, Rank: 3
Team: Clippers, Conference: West, Group: West Group B, Rank: 3
Team: Lakers, Conference: West, Group: West Group B, Rank: 1
Team: Heat, Conference: East, Group: East Group C, Rank: 2
Team: Bucks, Conference: East, Group: East Group C, Rank: 3
Team: Timberwolves, Conference: West, Group: West Group A, Rank: 3
Team: Nets, Conference: East, Group: East Group B, Rank: 5
Team: Knicks, Conference: East, Group: East Group C, Rank: 1
Team: Magic, Conference: East, Group: East Group B, Rank: 1
Team: Pacers, Conference: East, Group: East Group A, Rank: 4
Team: 76ers, Conference: East, Group: East Group B, Rank: 4
Team: Suns, Conference: West, Group: West Group A, Rank: 2
Team: Trail Blazers, Conference: West, Group: West Group C, Rank: 4
Team: Kings, Conference: West, Group: West Group A, Rank: 5
Team: Spurs, Conference: West, Group: West Group C, Rank: 1
Team: Thunder, Conference: West, Group: West Group A, Rank: 1
Team: Raptors, Conference: East, Group: East Group A, Rank: 1
Team: Jazz, Conference: West, Group: West Group A, Rank: 4
Team: Grizzlies, Conference: West, Group: West Group B, Rank: 2
Team: Wizards, Conference: East, Group: East Group A, Rank: 5
Team: Pistons, Conference: East, Group: East Group B, Rank: 3
Team: Hornets, Conference: East, Group: East Group C, Rank: 4

## 📜 Details

### Arg

```go title="arg.go"
type IstStandingsParams struct {
	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "2025-26"
	Season string `url:"Season"`

	// optional default "group"
	// "group" or "wildcard"
	Section string `url:"Section"`
}
```

### Response

[`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/response_ist_standings.go)
