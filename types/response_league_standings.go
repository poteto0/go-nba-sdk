package types

import "github.com/moznion/go-optional"

type LeagueStandingsRecord struct {
	LeagueID                string                  `json:"leagueID"`
	SeasonID                string                  `json:"seasonID"`
	TeamID                  int                     `json:"teamID"`
	TeamCity                string                  `json:"teamCity"`
	TeamName                string                  `json:"teamName"`
	Conference              string                  `json:"conference"`
	ConferenceRecord        string                  `json:"conferenceRecord"`
	PlayoffRank             int                     `json:"playoffRank"`
	ClinchIndicator         string                  `json:"clinchIndicator"`
	Division                string                  `json:"division"`
	DivisionRecord          string                  `json:"divisionRecord"`
	DivisionRank            int                     `json:"divisionRank"`
	Wins                    int                     `json:"wins"`
	Losses                  int                     `json:"losses"`
	WinPct                  float64                 `json:"winPct"`
	LeagueRank              int                     `json:"leagueRank"`
	Record                  string                  `json:"record"`
	Home                    string                  `json:"home"`
	Road                    string                  `json:"road"`
	L10                     string                  `json:"l10"`
	Last10Home              string                  `json:"last10Home"`
	Last10Road              string                  `json:"last10Road"`
	OT                      string                  `json:"ot"`
	ThreePTSOrLess          string                  `json:"threePTSOrLess"`
	TenPTSOrMore            string                  `json:"tenPTSOrMore"`
	LongHomeStreak          int                     `json:"longHomeStreak"`
	StrLongHomeStreak       string                  `json:"strLongHomeStreak"`
	LongRoadStreak          int                     `json:"longRoadStreak"`
	StrLongRoadStreak       string                  `json:"strLongRoadStreak"`
	LongWinStreak           int                     `json:"longWinStreak"`
	LongLossStreak          int                     `json:"longLossStreak"`
	CurrentHomeStreak       int                     `json:"currentHomeStreak"`
	StrCurrentHomeStreak    string                  `json:"strCurrentHomeStreak"`
	CurrentRoadStreak       int                     `json:"currentRoadStreak"`
	StrCurrentRoadStreak    string                  `json:"strCurrentRoadStreak"`
	CurrentStreak           int                     `json:"currentStreak"`
	StrCurrentStreak        string                  `json:"strCurrentStreak"`
	ConferenceGamesBack     float64                 `json:"conferenceGamesBack"`
	DivisionGamesBack       float64                 `json:"divisionGamesBack"`
	ClinchedConferenceTitle optional.Option[string] `json:"clinchedConferenceTitle"`
	ClinchedDivisionTitle   optional.Option[string] `json:"clinchedDivisionTitle"`
	ClinchedPlayoffBirth    optional.Option[string] `json:"clinchedPlayoffBirth"`
	EliminatedConference    optional.Option[string] `json:"eliminatedConference"`
	EliminatedDivision      optional.Option[string] `json:"eliminatedDivision"`
	AheadAtHalf             string                  `json:"aheadAtHalf"`
	BehindAtHalf            string                  `json:"behindAtHalf"`
	TiedAtHalf              string                  `json:"tiedAtHalf"`
	AheadAtThird            string                  `json:"aheadAtThird"`
	BehindAtThird           string                  `json:"behindAtThird"`
	TiedAtThird             string                  `json:"tiedAtThird"`
	Score100PTS             string                  `json:"score100PTS"`
	OppScore100PTS          string                  `json:"oppScore100PTS"`
	OppOver500              string                  `json:"oppOver500"`
	LeadInFGPCT             string                  `json:"leadInFGPCT"`
	LeadInReb               string                  `json:"leadInReb"`
	FewerTurnovers          string                  `json:"fewerTurnovers"`
	PointsPG                float64                 `json:"pointsPG"`
	OppPointsPG             float64                 `json:"oppPointsPG"`
	DiffPointsPG            float64                 `json:"diffPointsPG"`
	VsEast                  string                  `json:"vsEast"`
	VsAtlantic              string                  `json:"vsAtlantic"`
	VsCentral               string                  `json:"vsCentral"`
	VsSoutheast             string                  `json:"vsSoutheast"`
	VsWest                  string                  `json:"vsWest"`
	VsNorthwest             string                  `json:"vsNorthwest"`
	VsPacific               string                  `json:"vsPacific"`
	VsSouthwest             string                  `json:"vsSouthwest"`
	Jan                     optional.Option[string] `json:"jan"`
	Feb                     optional.Option[string] `json:"feb"`
	Mar                     optional.Option[string] `json:"mar"`
	Apr                     optional.Option[string] `json:"apr"`
	May                     optional.Option[string] `json:"may"`
	Jun                     optional.Option[string] `json:"jun"`
	Jul                     optional.Option[string] `json:"jul"`
	Aug                     optional.Option[string] `json:"aug"`
	Sep                     optional.Option[string] `json:"sep"`
	Oct                     optional.Option[string] `json:"oct"`
	Nov                     optional.Option[string] `json:"nov"`
	Dec                     optional.Option[string] `json:"dec"`
	PreAS                   optional.Option[string] `json:"preAS"`
	PostAS                  optional.Option[string] `json:"postAS"`
}

type LeagueStandingsResponseContent struct {
	SeasonYear string                  `json:"seasonYear"`
	Standings  []LeagueStandingsRecord `json:"standings"`
}
