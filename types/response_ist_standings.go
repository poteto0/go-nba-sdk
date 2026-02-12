package types

type IstStandingsResponseContent struct {
	LeagueID      string             `json:"leagueId"`
	SeasonYear    string             `json:"seasonYear"`
	UnixTimeStamp int64              `json:"unixTimeStamp"`
	TimeStampUtc  string             `json:"timeStampUtc"`
	Teams         []IstStandingsTeam `json:"teams"`
}

type IstStandingsTeam struct {
	TeamID              int                `json:"teamId"`
	TeamCity            string             `json:"teamCity"`
	TeamName            string             `json:"teamName"`
	TeamAbbreviation    string             `json:"teamAbbreviation"`
	TeamSlug            string             `json:"teamSlug"`
	Conference          string             `json:"conference"`
	IstGroup            string             `json:"istGroup"`
	ClinchIndicator     string             `json:"clinchIndicator"`
	ClinchedIstKnockout int                `json:"clinchedIstKnockout"`
	ClinchedIstGroup    int                `json:"clinchedIstGroup"`
	ClinchedIstWildcard int                `json:"clinchedIstWildcard"`
	IstWildcardRank     int                `json:"istWildcardRank"`
	IstGroupRank        int                `json:"istGroupRank"`
	IstKnockoutRank     *int               `json:"istKnockoutRank"`
	Wins                int                `json:"wins"`
	Losses              int                `json:"losses"`
	Pct                 float64            `json:"pct"`
	IstGroupGb          float64            `json:"istGroupGb"`
	IstWildcardGb       float64            `json:"istWildcardGb"`
	Diff                int                `json:"diff"`
	Pts                 int                `json:"pts"`
	OppPts              int                `json:"oppPts"`
	Games               []IstStandingsGame `json:"games"`
}

type IstStandingsGame struct {
	GameID                   string `json:"gameId"`
	GameNumber               int    `json:"gameNumber"`
	OpponentTeamAbbreviation string `json:"opponentTeamAbbreviation"`
	Location                 string `json:"location"`
	GameStatus               int    `json:"gameStatus"`
	GameStatusText           string `json:"gameStatusText"`
	Outcome                  string `json:"outcome"`
}
