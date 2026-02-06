---
sidebar_position: 1
slug: /
---

# go nba sdk

GoLang SDK for NBA API

### :art: Namespace API Structure

Intuitive to use

```go
client := gns.NewClient()I
// Stats Namespace
result := client.Stats.GetPlayerCareerStats(&types.PlayerCareerStatsParams{
	PlayerID: "203076",
})
```

### :zap: Quickly try out w/ golang

Can be used without being aware of the web API layer.

### :lock: nil-safe result

:::note

Dependent on the following library;

https://github.com/moznion/go-optional

:::

```go
client := gns.NewClient()
result := client.Stats.GetPlayerCareerStats(&types.PlayerCareerStatsParams{
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
```
