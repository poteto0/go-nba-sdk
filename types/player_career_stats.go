package types

import "github.com/moznion/go-optional"

type PlayerCareerStatsRecord struct {
	PlayerID        int                      `json:"playerId"`
	SeasonID        string                   `json:"seasonId"`
	LeagueID        string                   `json:"leagueId"`
	TeamID          int                      `json:"teamId"`
	TeamAbbeviation string                   `json:"teamAbbreviation"`
	PlayerAge       int                      `json:"playerAge"`
	GP              optional.Option[int]     `json:"gp"`
	GS              optional.Option[int]     `json:"gs"`
	Min             optional.Option[float64] `json:"min"`
	Pts             optional.Option[float64] `json:"pts"`
	FgM             optional.Option[float64] `json:"fgm"`
	FgA             optional.Option[float64] `json:"fga"`
	FgPct           optional.Option[float64] `json:"fgPct"`
	Fg3M            optional.Option[float64] `json:"fg3m"`
	Fg3A            optional.Option[float64] `json:"fg3a"`
	Fg3Pct          optional.Option[float64] `json:"fg3Pct"`
	FtM             optional.Option[float64] `json:"ftm"`
	FtA             optional.Option[float64] `json:"fta"`
	FtPct           optional.Option[float64] `json:"ftPct"`
	OReb            optional.Option[float64] `json:"oreb"`
	DReb            optional.Option[float64] `json:"dreb"`
	Reb             optional.Option[float64] `json:"reb"`
	Ast             optional.Option[float64] `json:"ast"`
	Stl             optional.Option[float64] `json:"stl"`
	Blk             optional.Option[float64] `json:"blk"`
	Tov             optional.Option[float64] `json:"tov"`
	PF              optional.Option[float64] `json:"pf"`
}

type PlayerCollegeSeasonStatsRecord struct {
	PlayerID       int                      `json:"playerId"`
	SeasonID       string                   `json:"seasonId"`
	LeagueID       string                   `json:"leagueId"`
	OrganizationID int                      `json:"organizationId"`
	SchoolName     string                   `json:"schoolName"`
	GP             optional.Option[int]     `json:"gp"`
	GS             optional.Option[int]     `json:"gs"`
	Min            optional.Option[float64] `json:"min"`
	Pts            optional.Option[float64] `json:"pts"`
	FgM            optional.Option[float64] `json:"fgm"`
	FgA            optional.Option[float64] `json:"fga"`
	FgPct          optional.Option[float64] `json:"fgPct"`
	Fg3M           optional.Option[float64] `json:"fg3m"`
	Fg3A           optional.Option[float64] `json:"fg3a"`
	Fg3Pct         optional.Option[float64] `json:"fg3Pct"`
	FtM            optional.Option[float64] `json:"ftm"`
	FtA            optional.Option[float64] `json:"fta"`
	FtPct          optional.Option[float64] `json:"ftPct"`
	OReb           optional.Option[float64] `json:"oreb"`
	DReb           optional.Option[float64] `json:"dreb"`
	Reb            optional.Option[float64] `json:"reb"`
	Ast            optional.Option[float64] `json:"ast"`
	Stl            optional.Option[float64] `json:"stl"`
	Blk            optional.Option[float64] `json:"blk"`
	Tov            optional.Option[float64] `json:"tov"`
	PF             optional.Option[float64] `json:"pf"`
}

type PlayerCollegeCareerStatsRecord struct {
	PlayerID       int                      `json:"playerId"`
	LeagueID       string                   `json:"leagueId"`
	OrganizationID int                      `json:"organizationId"`
	GP             optional.Option[int]     `json:"gp"`
	GS             optional.Option[int]     `json:"gs"`
	Min            optional.Option[float64] `json:"min"`
	Pts            optional.Option[float64] `json:"pts"`
	FgM            optional.Option[float64] `json:"fgm"`
	FgA            optional.Option[float64] `json:"fga"`
	FgPct          optional.Option[float64] `json:"fgPct"`
	Fg3M           optional.Option[float64] `json:"fg3m"`
	Fg3A           optional.Option[float64] `json:"fg3a"`
	Fg3Pct         optional.Option[float64] `json:"fg3Pct"`
	FtM            optional.Option[float64] `json:"ftm"`
	FtA            optional.Option[float64] `json:"fta"`
	FtPct          optional.Option[float64] `json:"ftPct"`
	OReb           optional.Option[float64] `json:"oreb"`
	DReb           optional.Option[float64] `json:"dreb"`
	Reb            optional.Option[float64] `json:"reb"`
	Ast            optional.Option[float64] `json:"ast"`
	Stl            optional.Option[float64] `json:"stl"`
	Blk            optional.Option[float64] `json:"blk"`
	Tov            optional.Option[float64] `json:"tov"`
	PF             optional.Option[float64] `json:"pf"`
}

type PlayerCareerRankingsRecord struct {
	PlayerID        int    `json:"playerId"`
	SeasonID        string `json:"seasonId"`
	LeagueID        string `json:"leagueId"`
	TeamID          int    `json:"teamId"`
	TeamAbbeviation string `json:"teamAbbreviation"`
	/* means nothing
	PlayerAge       int    `json:"playerAge"`
	GP              int    `json:"gp"`
	GS              int    `json:"gs"`
	*/
	RankPgMin    optional.Option[int] `json:"rankPgMin"`
	RankPgPts    optional.Option[int] `json:"rankPgPts"`
	RankPgOReb   optional.Option[int] `json:"rankPgOreb"`
	RankPgDReb   optional.Option[int] `json:"rankPgDreb"`
	RankPgReb    optional.Option[int] `json:"rankPgReb"`
	RankPgAst    optional.Option[int] `json:"rankPgAst"`
	RankPgStl    optional.Option[int] `json:"rankPgStl"`
	RankPgBlk    optional.Option[int] `json:"rankPgBlk"`
	RankPgTov    optional.Option[int] `json:"rankPgTov"`
	RankPgFgm    optional.Option[int] `json:"rankPgFgm"`
	RankPgFga    optional.Option[int] `json:"rankPgFga"`
	RankPgFgPct  optional.Option[int] `json:"rankPgFgPct"`
	RankPgFg3m   optional.Option[int] `json:"rankPgFg3m"`
	RankPgFg3a   optional.Option[int] `json:"rankPgFg3a"`
	RankPgFg3Pct optional.Option[int] `json:"rankPgFg3Pct"`
	RankPgFtm    optional.Option[int] `json:"rankPgFtm"`
	RankPgFta    optional.Option[int] `json:"rankPgFta"`
	RankPgFtPct  optional.Option[int] `json:"rankPgFtPct"`
	RankPgEff    optional.Option[int] `json:"rankPgEff"`
}

type PlayerHighsRecord struct {
	PlayerID           int     `json:"playerId"`
	GameID             string  `json:"gameId"`
	GameDate           string  `json:"gameDate"`
	VsTeamID           int     `json:"vsTeamId"`
	VsTeamCity         string  `json:"vsTeamCity"`
	VsTeamName         string  `json:"vsTeamName"`
	VsTeamAbbreviation string  `json:"vsTeamAbbreviation"`
	Stat               string  `json:"stat"`
	StatValue          float64 `json:"statValue"`
	StatOrder          int     `json:"statOrder"`
	DateEst            string  `json:"dateEst"`
}

type PlayerCareerStatsResponseContent struct {
	// nba stats
	SeasonTotalsRegularSeason  []PlayerCareerStatsRecord `json:"seasonTotalsRegularSeason"`
	CareerTotalsRegularSeason  []PlayerCareerStatsRecord `json:"careerTotalsRegularSeason"`
	SeasonTotalsPostSeason     []PlayerCareerStatsRecord `json:"seasonTotalsPostSeason"`
	CareerTotalsPostSeason     []PlayerCareerStatsRecord `json:"careerTotalsPostSeason"`
	SeasonTotalsAllStarSeason  []PlayerCareerStatsRecord `json:"seasonTotalsAllStarSeason"`
	CareerTotalsAllStarSeason  []PlayerCareerStatsRecord `json:"careerTotalsAllStarSeason"`
	SeasonTotalsShowcaseSeason []PlayerCareerStatsRecord `json:"seasonTotalsShowcaseSeason"`
	CareerTotalsShowcaseSeason []PlayerCareerStatsRecord `json:"careerTotalsShowcaseSeason"`

	// college stats
	SeasonTotalsCollegeSeason []PlayerCollegeSeasonStatsRecord `json:"seasonTotalsCollegeSeason"`
	CareerTotalsCollegeSeason []PlayerCollegeSeasonStatsRecord `json:"careerTotalsCollegeSeason"`

	// rankings
	SeasonRankingsRegularSeason []PlayerCareerRankingsRecord `json:"seasonRankingsRegularSeason"`
	SeasonRankingsPostSeason    []PlayerCareerRankingsRecord `json:"seasonRankingsPostSeason"`

	// highs
	SeasonHighs []PlayerHighsRecord `json:"seasonHighs"`
	CareerHighs []PlayerHighsRecord `json:"careerHighs"`
}
