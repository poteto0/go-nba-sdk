package types

type LivePlayByPlayResponse struct {
	Meta Meta           `json:"meta"`
	Game PlayByPlayGame `json:"game"`
}

type PlayByPlayGame struct {
	GameID  string   `json:"gameId"`
	Actions []Action `json:"actions"`
}

type Action struct {
	ActionNumber            int      `json:"actionNumber"`
	ActionType              string   `json:"actionType"`
	AssistPersonID          int      `json:"assistPersonId,omitempty"`
	AssistPlayerNameInitial string   `json:"assistPlayerNameInitial,omitempty"`
	AssistTotal             int      `json:"assistTotal,omitempty"`
	BlockPersonID           int      `json:"blockPersonId,omitempty"`
	BlockPlayerName         string   `json:"blockPlayerName,omitempty"`
	Clock                   string   `json:"clock"`
	Description             string   `json:"description,omitempty"`
	Descriptor              string   `json:"descriptor,omitempty"`
	Edited                  string   `json:"edited,omitempty"`
	FoulDrawnPersonID       int      `json:"foulDrawnPersonId,omitempty"`
	FoulDrawnPlayerName     string   `json:"foulDrawnPlayerName,omitempty"`
	FoulPersonalTotal       int      `json:"foulPersonalTotal,omitempty"`
	FoulTechnicalTotal      int      `json:"foulTechnicalTotal,omitempty"`
	IsFieldGoal             int      `json:"isFieldGoal,omitempty"`
	JumpBallLostPersonID    int      `json:"jumpBallLostPersonId,omitempty"`
	JumpBallLostPlayerName  string   `json:"jumpBallLostPlayerName,omitempty"`
	JumpBallRecoveredName   string   `json:"jumpBallRecoveredName,omitempty"`
	JumpBallWonPersonID     int      `json:"jumpBallWonPersonId,omitempty"`
	JumpBallWonPlayerName   string   `json:"jumpBallWonPlayerName,omitempty"`
	OfficialID              int      `json:"officialId,omitempty"`
	OrderNumber             int      `json:"orderNumber,omitempty"`
	Period                  int      `json:"period"`
	PeriodType              string   `json:"periodType,omitempty"`
	PersonID                int      `json:"personId,omitempty"`
	PersonIdsFilter         []int    `json:"personIdsFilter,omitempty"`
	PlayerName              string   `json:"playerName,omitempty"`
	PlayerNameI             string   `json:"playerNameI,omitempty"`
	PointsTotal             int      `json:"pointsTotal,omitempty"`
	Possession              int      `json:"possession,omitempty"`
	Qualifiers              []string `json:"qualifiers,omitempty"`
	ReboundDefensiveTotal   int      `json:"reboundDefensiveTotal,omitempty"`
	ReboundOffensiveTotal   int      `json:"reboundOffensiveTotal,omitempty"`
	ReboundTotal            int      `json:"reboundTotal,omitempty"`
	ScoreAway               string   `json:"scoreAway,omitempty"`
	ScoreHome               string   `json:"scoreHome,omitempty"`
	ShotActionNumber        int      `json:"shotActionNumber,omitempty"`
	ShotDistance            float64  `json:"shotDistance,omitempty"`
	ShotResult              string   `json:"shotResult,omitempty"`
	Side                    string   `json:"side,omitempty"`
	StealPersonID           int      `json:"stealPersonId,omitempty"`
	StealPlayerName         string   `json:"stealPlayerName,omitempty"`
	SubType                 string   `json:"subType,omitempty"`
	TeamID                  int      `json:"teamId,omitempty"`
	TeamTricode             string   `json:"teamTricode,omitempty"`
	TimeActual              string   `json:"timeActual,omitempty"`
	TurnoverTotal           int      `json:"turnoverTotal,omitempty"`
	X                       float64  `json:"x,omitempty"`
	XLegacy                 int      `json:"xLegacy,omitempty"`
	Y                       float64  `json:"y,omitempty"`
	YLegacy                 int      `json:"yLegacy,omitempty"`
}
