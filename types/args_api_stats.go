package types

type PlayerCareerStatsParams struct {
	// !required
	PlayerID string `url:"PlayerID"`

	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "PerGame"
	PerMode string `url:"PerMode"`
}
