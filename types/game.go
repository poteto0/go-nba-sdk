package types

import (
	"strings"
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
func (g *Game) IsFinished() bool {
	return g.GameStatus == 3
}

// parse clock
// PT07M11.01S -> 07:11.01
func (g *Game) Clock() string {
	return parseClockTypeStr(g.GameClock)
}

func (g *Game) ClockMinutes() (float64, bool) {
	minutesSplit := strings.Split(g.Clock(), ":")[0]
	return parseNumFromParsedClockTypeStr(minutesSplit)
}

func (g *Game) ClockSeconds() (float64, bool) {
	secondsSplit := strings.Split(g.Clock(), ":")[1]
	return parseNumFromParsedClockTypeStr(secondsSplit)
}
