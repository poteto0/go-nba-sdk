package types

type ScheduleLeagueV2Response struct {
	Meta           Meta           `json:"meta"`
	LeagueSchedule LeagueSchedule `json:"leagueSchedule"`
}

type LeagueSchedule struct {
	SeasonYear string     `json:"seasonYear"`
	LeagueID   string     `json:"leagueId"`
	GameDates  []GameDate `json:"gameDates"`
}

type GameDate struct {
	GameDate string         `json:"gameDate"`
	Games    []ScheduleGame `json:"games"`
}

type ScheduleGame struct {
	GameID           string         `json:"gameId"`
	GameCode         string         `json:"gameCode"`
	GameStatus       int            `json:"gameStatus"`
	GameStatusText   string         `json:"gameStatusText"`
	GameSequence     int            `json:"gameSequence"`
	GameDateEst      string         `json:"gameDateEst"`
	GameTimeEst      string         `json:"gameTimeEst"`
	GameDateTimeEst  string         `json:"gameDateTimeEst"`
	GameDateUTC      string         `json:"gameDateUTC"`
	GameTimeUTC      string         `json:"gameTimeUTC"`
	GameDateTimeUTC  string         `json:"gameDateTimeUTC"`
	AwayTeamTime     string         `json:"awayTeamTime"`
	HomeTeamTime     string         `json:"homeTeamTime"`
	Day              string         `json:"day"`
	MonthNum         int            `json:"monthNum"`
	WeekNumber       int            `json:"weekNumber"`
	WeekName         string         `json:"weekName"`
	IfNecessary      string         `json:"ifNecessary"` // "false"
	SeriesGameNumber string         `json:"seriesGameNumber"`
	GameLabel        string         `json:"gameLabel"`
	GameSubLabel     string         `json:"gameSubLabel"`
	SeriesText       string         `json:"seriesText"`
	ArenaName        string         `json:"arenaName"`
	ArenaState       string         `json:"arenaState"`
	ArenaCity        string         `json:"arenaCity"`
	PostponedStatus  string         `json:"postponedStatus"`
	BranchLink       string         `json:"branchLink"`
	GameSubtype      string         `json:"gameSubtype"`
	IsNeutral        bool           `json:"isNeutral"`
	Broadcasters     Broadcasters   `json:"broadcasters"`
	HomeTeam         ScheduleTeam   `json:"homeTeam"`
	AwayTeam         ScheduleTeam   `json:"awayTeam"`
	PointsLeaders    []PointsLeader `json:"pointsLeaders"`
}

type Broadcasters struct {
	NationalBroadcasters      []Broadcaster `json:"nationalBroadcasters"`
	NationalRadioBroadcasters []Broadcaster `json:"nationalRadioBroadcasters"`
	NationalOttBroadcasters   []Broadcaster `json:"nationalOttBroadcasters"`
	HomeTvBroadcasters        []Broadcaster `json:"homeTvBroadcasters"`
	HomeRadioBroadcasters     []Broadcaster `json:"homeRadioBroadcasters"`
	HomeOttBroadcasters       []Broadcaster `json:"homeOttBroadcasters"`
	AwayTvBroadcasters        []Broadcaster `json:"awayTvBroadcasters"`
	AwayRadioBroadcasters     []Broadcaster `json:"awayRadioBroadcasters"`
	AwayOttBroadcasters       []Broadcaster `json:"awayOttBroadcasters"`
}

type Broadcaster struct {
	BroadcasterScope        string `json:"broadcasterScope"`
	BroadcasterMedia        string `json:"broadcasterMedia"`
	BroadcasterID           int    `json:"broadcasterId"`
	BroadcasterDisplay      string `json:"broadcasterDisplay"`
	BroadcasterAbbreviation string `json:"broadcasterAbbreviation"`
	BroadcasterDescription  string `json:"broadcasterDescription"`
	TapeDelayComments       string `json:"tapeDelayComments"`
	BroadcasterVideoLink    string `json:"broadcasterVideoLink"`
	BroadcasterTeamID       int    `json:"broadcasterTeamId"`
	BroadcasterRanking      *int   `json:"broadcasterRanking"`
	LocalizationRegion      string `json:"localizationRegion"`
}

type ScheduleTeam struct {
	TeamID      int    `json:"teamId"`
	TeamName    string `json:"teamName"`
	TeamCity    string `json:"teamCity"`
	TeamTricode string `json:"teamTricode"`
	TeamSlug    string `json:"teamSlug"`
	Wins        int    `json:"wins"`
	Losses      int    `json:"losses"`
	Score       int    `json:"score"`
	Seed        int    `json:"seed"` // Sometimes null? assume int for now
}

type PointsLeader struct {
	PersonID    int     `json:"personId"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	TeamID      int     `json:"teamId"`
	TeamCity    string  `json:"teamCity"`
	TeamName    string  `json:"teamName"`
	TeamTricode string  `json:"teamTricode"`
	Points      float64 `json:"points"`
}
