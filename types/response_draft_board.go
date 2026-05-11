package types

type DraftBoardParams struct {
	// default 00, 00 for NBA, 20 for WNBA
	LeagueID string `url:"LeagueID"`

	// !required
	SeasonYear string `url:"SeasonYear"`

	College     string `url:"College,omitempty"`
	OverallPick string `url:"OverallPick,omitempty"`
	RoundNum    string `url:"RoundNum,omitempty"`
	RoundPick   string `url:"RoundPick,omitempty"`
	TeamID      string `url:"TeamID,omitempty"`
	TopX        string `url:"TopX,omitempty"`
}

type DraftBoardResponse struct {
	Resource   *string `json:"resource,omitempty"`
	Parameters *struct {
		LeagueID *string `json:"LeagueID,omitempty"`
		Season   *string `json:"Season,omitempty"`
	} `json:"parameters,omitempty"`
	ResultSets []DraftBoardResultSet `json:"resultSets,omitempty"`
}

type DraftBoardResultSet struct {
	Name    *string  `json:"name,omitempty"`
	Headers []string `json:"headers,omitempty"`
	RowSet  [][]any  `json:"rowSet,omitempty"`
}
