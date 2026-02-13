package types

type PlayerCareerStatsParams struct {
	// !required
	PlayerID string `url:"PlayerID"`

	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "PerGame"
	PerMode string `url:"PerMode"`
}

type IstStandingsParams struct {
	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "2025-26"
	Season string `url:"Season"`

	// optional default "group"
	// group or wildcard
	Section string `url:"Section"`
}

type LeagueStandingsParams struct {
	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "2025-26"
	Season string `url:"Season"`

	// optional default "Regular Season"
	// "Regular Season" or "Pre Season"
	SeasonType string `url:"SeasonType"`

	// optional
	SeasonYear string `url:"SeasonYear"`
}

type ScheduleLeagueV2Params struct {
	// optional default "00"
	LeagueID string `url:"LeagueID"`

	// optional default "2025-26"
	Season string `url:"Season"`
}
