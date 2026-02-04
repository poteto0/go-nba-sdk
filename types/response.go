package types

type Response[T any] struct {
	Contents   T     `json:"contents"`
	StatusCode int   `json:"status_code"`
	Error      error `json:"error"`
}

type ResultSet struct {
	Name    string   `json:"name"`
	Headers []string `json:"headers"`
	RowSet  [][]any  `json:"rowSet"`
}

type RawResponse struct {
	Resource   string         `json:"resource"`
	Parameters map[string]any `json:"parameters"`
	ResultSets []ResultSet    `json:"resultSets"`
}
