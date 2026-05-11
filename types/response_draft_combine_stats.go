package types

type DraftCombineStatsResponseContent struct {
	CombineStats []DraftCombineStatsRecord `json:"combineStats"`
}

type DraftCombineStatsRecord struct {
	Season              string   `json:"season"`
	PlayerID            int      `json:"playerId"`
	FirstName           string   `json:"firstName"`
	LastName            string   `json:"lastName"`
	PlayerName          string   `json:"playerName"`
	Position            string   `json:"position"`
	HeightWoShoes       *float64 `json:"heightWoShoes"`
	HeightWShoes        *float64 `json:"heightWShoes"`
	Weight              *float64 `json:"weight"`
	Wingspan            *float64 `json:"wingspan"`
	StandingReach       *float64 `json:"standingReach"`
	BodyFatPct          *float64 `json:"bodyFatPct"`
	HandLength          *float64 `json:"handLength"`
	HandWidth           *float64 `json:"handWidth"`
	StandingVertical    *float64 `json:"standingVerticalLeap"`
	MaxVertical         *float64 `json:"maxVerticalLeap"`
	LaneAgility         *float64 `json:"laneAgilityTime"`
	ModifiedLaneAgility *float64 `json:"modifiedLaneAgilityTime"`
	ThreeQuarterSprint  *float64 `json:"threeQuarterSprint"`
	BenchPress          *int     `json:"benchPress"`
}
