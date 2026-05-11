package types

type DraftCombineStatsParams struct {
	LeagueID   string `url:"LeagueID,omitempty"`
	SeasonYear string `url:"SeasonYear,omitempty"`
}

type DraftCombineStatsResponse struct {
	Resource   *string      `json:"resource"`
	Parameters *interface{} `json:"parameters"`
	ResultSets []struct {
		Name    string          `json:"name"`
		Headers []string        `json:"headers"`
		RowSet  [][]interface{} `json:"rowSet"`
	} `json:"resultSets"`
}
