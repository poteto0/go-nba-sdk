package types

import (
	"time"

	"github.com/moznion/go-optional"
)

type Game struct {
	GameId            string                      `json:"gameId"`
	GameCode          string                      `json:"gameCode"`
	GameStatus        int                         `json:"gameStatus"`
	GameStatusText    string                      `json:"gameStatusText"`
	Period            int                         `json:"period"`
	GameClock         string                      `json:"gameClock"`
	GameTimeUTC       time.Time                   `json:"gameTimeUTC"`
	GameEt            time.Time                   `json:"gameEt"`
	RegulationPeriods int                         `json:"regulationPeriods"`
	IfNecessary       bool                        `json:"ifNecessary"`
	SeriesGameNumber  string                      `json:"seriesGameNumber"`
	GameLabel         string                      `json:"gameLabel"`
	GameSubLabel      string                      `json:"gameSubLabel"`
	SeriesText        string                      `json:"seriesText"`
	SeriesConference  string                      `json:"seriesConference"`
	PoRoundDesc       string                      `json:"poRoundDesc"`
	GameSubtype       string                      `json:"gameSubtype"`
	IsNeutral         bool                        `json:"isNeutral"`
	HomeTeam          Team                        `json:"homeTeam"`
	AwayTeam          Team                        `json:"awayTeam"`
	GameLeaders       GameLeaders                 `json:"gameLeaders"`
	PbOdds            PbOdds                      `json:"pbOdds"`
	Arena             optional.Option[Arena]      `json:"arena"`
	Officials         optional.Option[[]Official] `json:"officials"`
}

// if the game is finished or not
func (g Game) IsFinished() bool {
	return g.GameStatus == 3
}
