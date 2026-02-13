package types

type PlayerCareerStatsRecord struct {
	PlayerID        int      `json:"playerId"`
	SeasonID        string   `json:"seasonId"`
	LeagueID        string   `json:"leagueId"`
	TeamID          int      `json:"teamId"`
	TeamAbbeviation string   `json:"teamAbbreviation"`
	PlayerAge       int      `json:"playerAge"`
	GP              *int     `json:"gp"`
	GS              *int     `json:"gs"`
	Min             *float64 `json:"min"`
	Pts             *float64 `json:"pts"`
	FgM             *float64 `json:"fgm"`
	FgA             *float64 `json:"fga"`
	FgPct           *float64 `json:"fgPct"`
	Fg3M            *float64 `json:"fg3m"`
	Fg3A            *float64 `json:"fg3a"`
	Fg3Pct          *float64 `json:"fg3Pct"`
	FtM             *float64 `json:"ftm"`
	FtA             *float64 `json:"fta"`
	FtPct           *float64 `json:"ftPct"`
	OReb            *float64 `json:"oreb"`
	DReb            *float64 `json:"dreb"`
	Reb             *float64 `json:"reb"`
	Ast             *float64 `json:"ast"`
	Stl             *float64 `json:"stl"`
	Blk             *float64 `json:"blk"`
	Tov             *float64 `json:"tov"`
	PF              *float64 `json:"pf"`
}

type PlayerCollegeSeasonStatsRecord struct {
	PlayerID       int      `json:"playerId"`
	SeasonID       string   `json:"seasonId"`
	LeagueID       string   `json:"leagueId"`
	OrganizationID int      `json:"organizationId"`
	SchoolName     string   `json:"schoolName"`
	GP             *int     `json:"gp"`
	GS             *int     `json:"gs"`
	Min            *float64 `json:"min"`
	Pts            *float64 `json:"pts"`
	FgM            *float64 `json:"fgm"`
	FgA            *float64 `json:"fga"`
	FgPct          *float64 `json:"fgPct"`
	Fg3M           *float64 `json:"fg3m"`
	Fg3A           *float64 `json:"fg3a"`
	Fg3Pct         *float64 `json:"fg3Pct"`
	FtM            *float64 `json:"ftm"`
	FtA            *float64 `json:"fta"`
	FtPct          *float64 `json:"ftPct"`
	OReb           *float64 `json:"oreb"`
	DReb           *float64 `json:"dreb"`
	Reb            *float64 `json:"reb"`
	Ast            *float64 `json:"ast"`
	Stl            *float64 `json:"stl"`
	Blk            *float64 `json:"blk"`
	Tov            *float64 `json:"tov"`
	PF             *float64 `json:"pf"`
}

type PlayerCollegeCareerStatsRecord struct {
	PlayerID       int      `json:"playerId"`
	LeagueID       string   `json:"leagueId"`
	OrganizationID int      `json:"organizationId"`
	GP             *int     `json:"gp"`
	GS             *int     `json:"gs"`
	Min            *float64 `json:"min"`
	Pts            *float64 `json:"pts"`
	FgM            *float64 `json:"fgm"`
	FgA            *float64 `json:"fga"`
	FgPct          *float64 `json:"fgPct"`
	Fg3M           *float64 `json:"fg3m"`
	Fg3A           *float64 `json:"fg3a"`
	Fg3Pct         *float64 `json:"fg3Pct"`
	FtM            *float64 `json:"ftm"`
	FtA            *float64 `json:"fta"`
	FtPct          *float64 `json:"ftPct"`
	OReb           *float64 `json:"oreb"`
	DReb           *float64 `json:"dreb"`
	Reb            *float64 `json:"reb"`
	Ast            *float64 `json:"ast"`
	Stl            *float64 `json:"stl"`
	Blk            *float64 `json:"blk"`
	Tov            *float64 `json:"tov"`
	PF             *float64 `json:"pf"`
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
	RankPgMin    *int `json:"rankPgMin"`
	RankPgPts    *int `json:"rankPgPts"`
	RankPgOReb   *int `json:"rankPgOreb"`
	RankPgDReb   *int `json:"rankPgDreb"`
	RankPgReb    *int `json:"rankPgReb"`
	RankPgAst    *int `json:"rankPgAst"`
	RankPgStl    *int `json:"rankPgStl"`
	RankPgBlk    *int `json:"rankPgBlk"`
	RankPgTov    *int `json:"rankPgTov"`
	RankPgFgm    *int `json:"rankPgFgm"`
	RankPgFga    *int `json:"rankPgFga"`
	RankPgFgPct  *int `json:"rankPgFgPct"`
	RankPgFg3m   *int `json:"rankPgFg3m"`
	RankPgFg3a   *int `json:"rankPgFg3a"`
	RankPgFg3Pct *int `json:"rankPgFg3Pct"`
	RankPgFtm    *int `json:"rankPgFtm"`
	RankPgFta    *int `json:"rankPgFta"`
	RankPgFtPct  *int `json:"rankPgFtPct"`
	RankPgEff    *int `json:"rankPgEff"`
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
