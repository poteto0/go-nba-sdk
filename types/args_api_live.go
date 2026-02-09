package types

type ScoreBoardParams struct {
	// optional default "00"
	LeagueID string
}

type BoxScoreParams struct {
	// !required
	GameID string
}
