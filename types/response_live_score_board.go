package types

import (
	"github.com/moznion/go-optional"
)

type LiveScoreBoardResponse struct {
	Meta       Meta       `json:"meta"`
	Scoreboard Scoreboard `json:"scoreboard"`
}

type Scoreboard struct {
	GameDate   string `json:"gameDate"`
	LeagueId   string `json:"leagueId"`
	LeagueName string `json:"leagueName"`
	Games      []Game `json:"games"`
}

type Team struct {
	TeamId            int                                    `json:"teamId"`
	TeamName          string                                 `json:"teamName"`
	TeamCity          string                                 `json:"teamCity"`
	TeamTricode       string                                 `json:"teamTricode"`
	Wins              int                                    `json:"wins"`
	Losses            int                                    `json:"losses"`
	Score             int                                    `json:"score"`
	Seed              optional.Option[int]                   `json:"seed"`
	InBonus           optional.Option[string]                `json:"inBonus"`
	TimeoutsRemaining int                                    `json:"timeoutsRemaining"`
	Periods           []Period                               `json:"periods"`
	Players           optional.Option[[]Player]              `json:"players"`
	Statistics        optional.Option[TeamBoxScoreStatistic] `json:"statistics"`
}

type Period struct {
	Period     int    `json:"period"`
	PeriodType string `json:"periodType"`
	Score      int    `json:"score"`
}

type GameLeaders struct {
	HomeLeaders Leader `json:"homeLeaders"`
	AwayLeaders Leader `json:"awayLeaders"`
}

type Leader struct {
	PersonId    int                     `json:"personId"`
	Name        string                  `json:"name"`
	JerseyNum   string                  `json:"jerseyNum"`
	Position    string                  `json:"position"`
	TeamTricode string                  `json:"teamTricode"`
	PlayerSlug  optional.Option[string] `json:"playerSlug"`
	Points      int                     `json:"points"`
	Rebounds    int                     `json:"rebounds"`
	Assists     int                     `json:"assists"`
}

type PbOdds struct {
	Team      optional.Option[string] `json:"team"`
	Odds      float64                 `json:"odds"`
	Suspended int                     `json:"suspended"`
}
