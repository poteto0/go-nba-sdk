---
sidebar_position: 2
---

# Combine Stats

get Draft Combine Stats

## ⚡ Quick Start

```go title="main.go"
func main() {
	client := gns.NewClient(nil)
	result := client.Draft.GetCombineStats(&types.DraftCombineStatsParams{
		LeagueID:   "00",
		SeasonYear: "2025-26",
	})

	if result.Error != nil {
		panic(result.Error)
	}

	for _, resultSet := range resCombine.Contents.ResultSets {
		for _, record := range resultSet.CombineStats {
			var standingVert float64
			if record.StandingVertical != nil {
				standingVert = *record.StandingVertical
			}

			fmt.Printf("Player: %s, Position: %s, Vertical: %v\n",
				record.PlayerName,
				record.Position,
				standingVert,
			)
		}
	}
}
```

## 📜 Details

### Arg

```go title="arg.go"
type DraftCombineStatsParams struct {
	// optional default "00"
	LeagueID string `url:"LeagueID,omitempty"`

	// required
	SeasonYear string `url:"SeasonYear,omitempty"`
}
```

### Response

[`response structure`](https://github.com/poteto0/go-nba-sdk/tree/main/types/response_draft_combine_stats.go)

- Nullable fields are represented as pointers.

  ```go
  result := draft.GetCombineStats(client, types.DraftCombineStatsParams{...})

  for _, resultSet := range result.Contents.ResultSets {
  	for _, record := range resultSet.RowSet {
  		// Example of handling nullable field
  		if record.StandingVertical != nil {
  			fmt.Printf("Vertical: %v\n", *record.StandingVertical)
  		}
  	}
  }
  ```
