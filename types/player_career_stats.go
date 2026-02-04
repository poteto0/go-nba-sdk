package types

type PlayerCareerStatsRecord struct {
	PlayerID        int     `json:"playerId"`
	SeasonID        string  `json:"seasonId"`
	LeagueID        string  `json:"leagueId"`
	TeamID          int     `json:"teamId"`
	TeamAbbeviation string  `json:"teamAbbreviation"`
	PlayerAge       int     `json:"playerAge"`
	GP              int     `json:"gp"`
	GS              int     `json:"gs"`
	MIN             float64 `json:"min"`
	PTS             float64 `json:"pts"`
	FGM             float64 `json:"fgm"`
	FGA             float64 `json:"fga"`
	FG_PCT          float64 `json:"fgPct"`
	FG3M            float64 `json:"fg3m"`
	FG3A            float64 `json:"fg3a"`
	FG3_PCT         float64 `json:"fg3Pct"`
	FTM             float64 `json:"ftm"`
	FTA             float64 `json:"fta"`
	FT_PCT          float64 `json:"ftPct"`
	OREB            float64 `json:"oreb"`
	DREB            float64 `json:"dreb"`
	REB             float64 `json:"reb"`
	AST             float64 `json:"ast"`
	STL             float64 `json:"stl"`
	BLK             float64 `json:"blk"`
	TOV             float64 `json:"tov"`
	PF              float64 `json:"pf"`
}

type PlayerCollegeStatsRecord struct {
	PlayerID        int     `json:"playerId"`
	SeasonID        string  `json:"seasonId"`
	LeagueID        string  `json:"leagueId"`
	ORGANIZATION_ID float64 `json:"organizationId"`
	SCHOOL_NAME     string  `json:"schoolName"`
	GP              int     `json:"gp"`
	GS              int     `json:"gs"`
	MIN             float64 `json:"min"`
	PTS             float64 `json:"pts"`
	FGM             float64 `json:"fgm"`
	FGA             float64 `json:"fga"`
	FG_PCT          float64 `json:"fgPct"`
	FG3M            float64 `json:"fg3m"`
	FG3A            float64 `json:"fg3a"`
	FG3_PCT         float64 `json:"fg3Pct"`
	FTM             float64 `json:"ftm"`
	FTA             float64 `json:"fta"`
	FT_PCT          float64 `json:"ftPct"`
	OREB            float64 `json:"oreb"`
	DREB            float64 `json:"dreb"`
	REB             float64 `json:"reb"`
	AST             float64 `json:"ast"`
	STL             float64 `json:"stl"`
	BLK             float64 `json:"blk"`
	TOV             float64 `json:"tov"`
	PF              float64 `json:"pf"`
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
	// rankings == 0 means no rank available
	RankPgMin    int `json:"rankPgMin"`
	RankPgPts    int `json:"rankPgPts"`
	RankPgOReb   int `json:"rankPgOreb"`
	RankPgDReb   int `json:"rankPgDreb"`
	RankPgReb    int `json:"rankPgReb"`
	RankPgAst    int `json:"rankPgAst"`
	RankPgStl    int `json:"rankPgStl"`
	RankPgBlk    int `json:"rankPgBlk"`
	RankPgTov    int `json:"rankPgTov"`
	RankPgFgm    int `json:"rankPgFgm"`
	RankPgFga    int `json:"rankPgFga"`
	RankPgFgPct  int `json:"rankPgFgPct"`
	RankPgFg3m   int `json:"rankPgFg3m"`
	RankPgFg3a   int `json:"rankPgFg3a"`
	RankPgFg3Pct int `json:"rankPgFg3Pct"`
	RankPgFtm    int `json:"rankPgFtm"`
	RankPgFta    int `json:"rankPgFta"`
	RankPgFtPct  int `json:"rankPgFtPct"`
	RankPgEff    int `json:"rankPgEff"`
}

type PlayerHighsRecord struct {
	PlayerID   int    `json:"playerId"`
	GameID     string `json:"gameId"`
	GameDate   string `json:"gameDate"`
	VsTeamId   int    `json:"vsTeamId"`
	VsTeamCity string `json:"vsTeamCity"`
	// TODO
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
	SeasonTotalsCollegeSeason []PlayerCollegeStatsRecord `json:"seasonTotalsCollegeSeason"`
	CareerTotalsCollegeSeason []PlayerCollegeStatsRecord `json:"careerTotalsCollegeSeason"`

	// rankings
	SeasonRankingsRegularSeason []PlayerCareerRankingsRecord `json:"seasonRankingsRegularSeason"`
	SeasonRankingsPostSeason    []PlayerCareerRankingsRecord `json:"seasonRankingsPostSeason"`
	//SeasonHighs []PlayerCareerStatsRecord `json:"seasonHighs"`
	//CareerHighs []PlayerCareerStatsRecord `json:"careerHighs"`
}
