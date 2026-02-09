package types

type Arena struct {
	ArenaId       int    `json:"arenaId"`
	ArenaName     string `json:"arenaName"`
	ArenaCity     string `json:"arenaCity"`
	ArenaState    string `json:"arenaState"`
	ArenaCountry  string `json:"arenaCountry"`
	ArenaTimezone string `json:"arenaTimezone"`
}
