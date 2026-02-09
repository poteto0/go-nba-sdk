package types

import "github.com/moznion/go-optional"

type CommonBoxScoreStatistic struct {
	Ast             optional.Option[int]     `json:"assists"`
	Blk             optional.Option[int]     `json:"blocks"`
	BlkReceived     optional.Option[int]     `json:"blocksReceived"`
	FgA             optional.Option[int]     `json:"fieldGoalsAttempted"`
	FgM             optional.Option[int]     `json:"fieldGoalsMade"`
	FgPct           optional.Option[float64] `json:"fieldGoalsPercentage"`
	PFOffensive     optional.Option[int]     `json:"foulsOffensive"`
	PFDrawn         optional.Option[int]     `json:"foulsDrawn"`
	PF              optional.Option[int]     `json:"foulsPersonal"`
	PFTechnical     optional.Option[int]     `json:"foulsTechnical"`
	FtA             optional.Option[int]     `json:"freeThrowsAttempted"`
	FtM             optional.Option[int]     `json:"freeThrowsMade"`
	FtPct           optional.Option[float64] `json:"freeThrowsPercentage"`
	Pts             optional.Option[int]     `json:"points"`
	PtsFastBreak    optional.Option[int]     `json:"pointsFastBreak"`
	PtsInPaint      optional.Option[int]     `json:"pointsInThePaint"`
	PtsSecondChance optional.Option[int]     `json:"pointsSecondChance"`
	DReb            optional.Option[int]     `json:"reboundsDefensive"`
	OReb            optional.Option[int]     `json:"reboundsOffensive"`
	Reb             optional.Option[int]     `json:"reboundsTotal"`
	Stl             optional.Option[int]     `json:"steals"`
	Fg3A            optional.Option[int]     `json:"threePointersAttempted"`
	Fg3M            optional.Option[int]     `json:"threePointersMade"`
	Fg3Pct          optional.Option[float64] `json:"threePointersPercentage"`
	Tov             optional.Option[int]     `json:"turnovers"`
	Fg2A            optional.Option[int]     `json:"twoPointersAttempted"`
	Fg2M            optional.Option[int]     `json:"twoPointersMade"`
	Fg2Pct          optional.Option[float64] `json:"twoPointersPercentage"`
}

type PlayerBoxScoreStatistic struct {
	CommonBoxScoreStatistic
	Minus             optional.Option[float64] `json:"minus"`
	Minutes           optional.Option[string]  `json:"minutes"`
	MinutesCalculated optional.Option[string]  `json:"minutesCalculated"`
	Plus              optional.Option[float64] `json:"plus"`
	PlusMinus         optional.Option[float64] `json:"plusMinusPoints"`
}

type TeamBoxScoreStatistic struct {
	CommonBoxScoreStatistic
	AstTovRatio            optional.Option[float64] `json:"assistTurnoverRatio"`
	BenchPts               optional.Option[int]     `json:"benchPoints"`
	BiggestLead            optional.Option[int]     `json:"biggestLead"`
	BiggestScoringRunScore optional.Option[string]  `json:"biggestScoringRunScore"`
	PtsFromTov             optional.Option[int]     `json:"pointsFromTurnovers"`
	PtsInPaintM            optional.Option[int]     `json:"pointsInThePaintMade"`
	PtsInPaintA            optional.Option[int]     `json:"pointsInThePaintAttempted"`
	PtsInPaintPct          optional.Option[float64] `json:"pointsInThePaintPercentage"`
	RebPersonal            optional.Option[int]     `json:"reboundsPersonal"`
	RebTeam                optional.Option[int]     `json:"reboundsTeam"`
	ORebTeam               optional.Option[int]     `json:"reboundsOffensiveTeam"`
	DRebTeam               optional.Option[int]     `json:"reboundsDefensiveTeam"`
	SecondChancePtsA       optional.Option[int]     `json:"secondChancePointsAttempted"`
	SecondChancePtsM       optional.Option[int]     `json:"secondChancePointsMade"`
	SecondChancePtsPct     optional.Option[float64] `json:"secondChancePointsPercentage"`
	TeamFGA                optional.Option[int]     `json:"teamFieldGoalsAttempted"`
	TimeLeading            optional.Option[string]  `json:"timeLeading"`
	TimesTied              optional.Option[int]     `json:"timesTied"`
	TSA                    optional.Option[float64] `json:"trueShootingAttempts"`
	TSM                    optional.Option[float64] `json:"trueShootingMade"`
	TSPct                  optional.Option[float64] `json:"trueShootingPercentage"`
	TovTeam                optional.Option[int]     `json:"turnoversTeam"`
	TovTotal               optional.Option[int]     `json:"turnoversTotal"`
}
