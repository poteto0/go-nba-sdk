package constants

// stats api
var (
	PlayerCareerStatsPath = "playercareerstats"
	IstStandingsPath      = "iststandings"
	LeagueStandingsPath   = "leaguestandings"
	ScheduleLeagueV2Path  = "scheduleleaguev2"
)

type ScheduleLeagueV2Params struct {
	LeagueID string
	Season   string
}

// live api
var (
	ScoreBoardPath   = "scoreboard"
	BoxScorePath     = "boxscore"
	PlayByPlayPath   = "playbyplay"
	PlayByPlayV3Path = "playbyplayv3"
)

// draft api
var (
	DraftCombineStatsPath = "draftcombinestats"
)
