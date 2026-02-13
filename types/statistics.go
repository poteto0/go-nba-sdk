package types

import (
	"strings"
)

type CommonBoxScoreStatistic struct {
	Minutes           string   `json:"minutes"`
	MinutesCalculated *string  `json:"minutesCalculated"`
	Ast               *int     `json:"assists"`
	Blk               *int     `json:"blocks"`
	BlkReceived       *int     `json:"blocksReceived"`
	FgA               *int     `json:"fieldGoalsAttempted"`
	FgM               *int     `json:"fieldGoalsMade"`
	FgPct             *float64 `json:"fieldGoalsPercentage"`
	PFOffensive       *int     `json:"foulsOffensive"`
	PFDrawn           *int     `json:"foulsDrawn"`
	PF                *int     `json:"foulsPersonal"`
	PFTechnical       *int     `json:"foulsTechnical"`
	FtA               *int     `json:"freeThrowsAttempted"`
	FtM               *int     `json:"freeThrowsMade"`
	FtPct             *float64 `json:"freeThrowsPercentage"`
	Pts               *int     `json:"points"`
	PtsFastBreak      *int     `json:"pointsFastBreak"`
	PtsInPaint        *int     `json:"pointsInThePaint"`
	PtsSecondChance   *int     `json:"pointsSecondChance"`
	DReb              *int     `json:"reboundsDefensive"`
	OReb              *int     `json:"reboundsOffensive"`
	Reb               *int     `json:"reboundsTotal"`
	Stl               *int     `json:"steals"`
	Fg3A              *int     `json:"threePointersAttempted"`
	Fg3M              *int     `json:"threePointersMade"`
	Fg3Pct            *float64 `json:"threePointersPercentage"`
	Tov               *int     `json:"turnovers"`
	Fg2A              *int     `json:"twoPointersAttempted"`
	Fg2M              *int     `json:"twoPointersMade"`
	Fg2Pct            *float64 `json:"twoPointersPercentage"`
}

// parse minutes
//
// PT07M11.01S -> 07:11.01
func (c *CommonBoxScoreStatistic) MinutesClock() string {
	return parseClockTypeStr(c.Minutes)
}

// parse & calculated minutes
//
// PT36M30.00S -> 36.5
func (c *CommonBoxScoreStatistic) Min() (float64, bool) {
	clockSplit := strings.Split(c.MinutesClock(), ":")
	if len(clockSplit) != 2 {
		return 0, false
	}

	minutesStr, secondsStr := clockSplit[0], clockSplit[1]
	minutes, ok := parseNumFromParsedClockTypeStr(minutesStr)
	if !ok {
		return 0, false
	}

	seconds, ok := parseNumFromParsedClockTypeStr(secondsStr)
	if !ok {
		return 0, false
	}

	return minutes + seconds/60, true
}

type PlayerBoxScoreStatistic struct {
	CommonBoxScoreStatistic
	Minus     *float64 `json:"minus"`
	Plus      *float64 `json:"plus"`
	PlusMinus *float64 `json:"plusMinusPoints"`
}

type TeamBoxScoreStatistic struct {
	CommonBoxScoreStatistic
	AstTovRatio            *float64 `json:"assistTurnoverRatio"`
	BenchPts               *int     `json:"benchPoints"`
	BiggestLead            *int     `json:"biggestLead"`
	BiggestScoringRunScore *string  `json:"biggestScoringRunScore"`
	PtsFromTov             *int     `json:"pointsFromTurnovers"`
	PtsInPaintM            *int     `json:"pointsInThePaintMade"`
	PtsInPaintA            *int     `json:"pointsInThePaintAttempted"`
	PtsInPaintPct          *float64 `json:"pointsInThePaintPercentage"`
	RebPersonal            *int     `json:"reboundsPersonal"`
	RebTeam                *int     `json:"reboundsTeam"`
	ORebTeam               *int     `json:"reboundsOffensiveTeam"`
	DRebTeam               *int     `json:"reboundsDefensiveTeam"`
	SecondChancePtsA       *int     `json:"secondChancePointsAttempted"`
	SecondChancePtsM       *int     `json:"secondChancePointsMade"`
	SecondChancePtsPct     *float64 `json:"secondChancePointsPercentage"`
	TeamFGA                *int     `json:"teamFieldGoalsAttempted"`
	TimeLeading            *string  `json:"timeLeading"`
	TimesTied              *int     `json:"timesTied"`
	TSA                    *float64 `json:"trueShootingAttempts"`
	TSM                    *float64 `json:"trueShootingMade"`
	TSPct                  *float64 `json:"trueShootingPercentage"`
	TovTeam                *int     `json:"turnoversTeam"`
	TovTotal               *int     `json:"turnoversTotal"`
}
