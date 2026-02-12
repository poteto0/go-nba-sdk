package types

type IstStandingsRecord struct {
	LeagueId             string  `json:"leagueId"`
	SeasonID             string  `json:"seasonId"`
	TeamID               int     `json:"teamId"`
	TeamCity             string  `json:"teamCity"`
	TeamName             string  `json:"teamName"`
	TeamAbbreviation     string  `json:"teamAbbreviation"`
	Conference           string  `json:"conference"`
	Group                string  `json:"group"`
	PlayoffRank          int     `json:"playoffRank"`
	W                    int     `json:"w"`
	L                    int     `json:"l"`
	WPct                 float64 `json:"wPct"`
	Home                 string  `json:"home"`
	Road                 string  `json:"road"`
	OT                   string  `json:"ot"`
	Last10               string  `json:"last10"`
	Streak               string  `json:"streak"`
	ConferenceRecord     string  `json:"conferenceRecord"`
	DivisionRecord       string  `json:"divisionRecord"`
	GroupRecord          string  `json:"groupRecord"`
	ClinchedPlayoff      int     `json:"clinchedPlayoff"`
	ClinchedPlayIn       int     `json:"clinchedPlayIn"`
	ClinchedConference   int     `json:"clinchedConference"`
	ClinchedDivision     int     `json:"clinchedDivision"`
	ClinchedGroup        int     `json:"clinchedGroup"`
	EliminatedPlayoff    int     `json:"eliminatedPlayoff"`
	EliminatedPlayIn     int     `json:"eliminatedPlayIn"`
	EliminatedConference int     `json:"eliminatedConference"`
	EliminatedDivision   int     `json:"eliminatedDivision"`
	EliminatedGroup      int     `json:"eliminatedGroup"`
}

type IstStandingsResponseContent struct {
	Standings []IstStandingsRecord `json:"standings"`
}
