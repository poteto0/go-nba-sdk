package types

import (
	"time"

	"github.com/moznion/go-optional"
)

type LiveScoreBoardResponse struct {
	Meta       Meta       `json:"meta"`
	Scoreboard Scoreboard `json:"scoreboard"`
}

type Meta struct {
	Version int    `json:"version"`
	Request string `json:"request"`
	Time    string `json:"time"`
	Code    int    `json:"code"`
}

type Scoreboard struct {
	GameDate   string `json:"gameDate"`
	LeagueId   string `json:"leagueId"`
	LeagueName string `json:"leagueName"`
	Games      []Game `json:"games"`
}

type Game struct {
	GameId            string      `json:"gameId"`
	GameCode          string      `json:"gameCode"`
	GameStatus        int         `json:"gameStatus"`
	GameStatusText    string      `json:"gameStatusText"`
	Period            int         `json:"period"`
	GameClock         string      `json:"gameClock"`
	GameTimeUTC       time.Time   `json:"gameTimeUTC"`
	GameEt            time.Time   `json:"gameEt"`
	RegulationPeriods int         `json:"regulationPeriods"`
	IfNecessary       bool        `json:"ifNecessary"`
	SeriesGameNumber  string      `json:"seriesGameNumber"`
	GameLabel         string      `json:"gameLabel"`
	GameSubLabel      string      `json:"gameSubLabel"`
	SeriesText        string      `json:"seriesText"`
	SeriesConference  string      `json:"seriesConference"`
	PoRoundDesc       string      `json:"poRoundDesc"`
	GameSubtype       string      `json:"gameSubtype"`
	IsNeutral         bool        `json:"isNeutral"`
	HomeTeam          Team        `json:"homeTeam"`
	AwayTeam          Team        `json:"awayTeam"`
	GameLeaders       GameLeaders `json:"gameLeaders"`
	PbOdds            PbOdds      `json:"pbOdds"`
}

// if the game is finished or not
func (g Game) IsFinished() bool {
	return g.GameStatus == 3
}

type Team struct {
	TeamId            int                   `json:"teamId"`
	TeamName          string                `json:"teamName"`
	TeamCity          string                `json:"teamCity"`
	TeamTricode       string                `json:"teamTricode"`
	Wins              int                   `json:"wins"`
	Losses            int                   `json:"losses"`
	Score             int                   `json:"score"`
	Seed              optional.Option[int]  `json:"seed"`
	InBonus           optional.Option[bool] `json:"inBonus"`
	TimeoutsRemaining int                   `json:"timeoutsRemaining"`
	Periods           []Period              `json:"periods"`
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
