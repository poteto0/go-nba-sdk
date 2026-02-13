---
sidebar_position: 1
slug: /
---

# go nba sdk

GoLang SDK for NBA API

### :art: Namespace API Structure

Intuitive to use

```go
client := gns.NewClient(nil)
// Stats Namespace
result := client.Stats.GetPlayerCareerStats(&types.PlayerCareerStatsParams{
	PlayerID: "203076",
})
```

### :zap: Quickly try out w/ golang

Can be used without being aware of the web API layer.

### :lock: Pointer-based optional fields

Nullable fields in response structs are represented as pointers.

```go
client := gns.NewClient(nil)
result := client.Stats.GetPlayerCareerStats(&types.PlayerCareerStatsParams{
	PlayerID: "203076",
})

fmt.Println("======================= RS ==========================")
for _, content := range result.Contents.SeasonTotalsRegularSeason {
	if content.Pts != nil {
		fmt.Printf("Season: %s, PPG: %.2f\n",
			content.SeasonID,
			*content.Pts,
		)
	}
}
```