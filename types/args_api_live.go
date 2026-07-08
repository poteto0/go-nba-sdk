package types

type ScoreBoardParams struct {
	// optional default "00"
	LeagueID string
}

type BoxScoreParams struct {
	// !required
	GameID string
}

type PlayByPlayParams struct {
	// !required
	GameID string

	// optional default is 1
	StartPeriod int

	// optional default is 4
	EndPeriod int
}
