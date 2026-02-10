package samples

var SampleLivePlayByPlayResponse = `
{
  "meta": {
    "version": 1,
    "code": 200,
    "request": "http://nba.cloud/games/0022000001/playbyplay?Format=json",
    "time": "2020-12-22 23:23:42.771674"
  },
  "game": {
    "gameId": "0022000001",
    "actions": [
      {
        "actionNumber": 2,
        "actionType": "period",
        "clock": "PT12M00.00S",
        "description": "Period Start",
        "period": 1,
        "periodType": "REGULAR",
        "subType": "start",
        "timeActual": "2020-12-23T00:06:19.0Z"
      },
      {
        "actionNumber": 4,
        "actionType": "jumpball",
        "clock": "PT11M58.00S",
        "description": "Jump Ball D. Jordan vs. J. Wiseman: Tip to Team (BKN)",
        "descriptor": "startperiod",
        "jumpBallLostPersonId": 1630164,
        "jumpBallLostPlayerName": "Wiseman",
        "jumpBallRecoveredName": "Team (BKN)",
        "jumpBallWonPersonId": 201599,
        "jumpBallWonPlayerName": "Jordan",
        "period": 1,
        "subType": "recovered",
        "teamId": 1610612751,
        "teamTricode": "BKN"
      },
      {
        "actionNumber": 12,
        "actionType": "2pt",
        "clock": "PT11M22.00S",
        "description": "K. Irving 21' pullup Jump Shot (2 PTS) (K. Durant 1 AST)",
        "isFieldGoal": 1,
        "period": 1,
        "personId": 202681,
        "playerName": "Irving",
        "playerNameI": "K. Irving",
        "pointsTotal": 2,
        "scoreAway": "2",
        "scoreHome": "2",
        "shotDistance": 21.97,
        "shotResult": "Made",
        "subType": "Jump Shot",
        "teamId": 1610612751,
        "teamTricode": "BKN",
        "x": 71.04139290407359,
        "y": 49.57873774509804
      },
      {
        "actionNumber": 16,
        "actionType": "3pt",
        "clock": "PT10M49.00S",
        "description": "K. Durant 25' 3PT pullup (3 PTS) (J. Harris 1 AST)",
        "isFieldGoal": 1,
        "period": 1,
        "personId": 201142,
        "playerName": "Durant",
        "playerNameI": "K. Durant",
        "pointsTotal": 3,
        "scoreAway": "4",
        "scoreHome": "5",
        "shotDistance": 25.88,
        "shotResult": "Made",
        "subType": "Jump Shot",
        "teamId": 1610612751,
        "teamTricode": "BKN"
      },
      {
        "actionNumber": 19,
        "actionType": "rebound",
        "clock": "PT10M25.00S",
        "description": "J. Harris REBOUND (Off:0 Def:1)",
        "period": 1,
        "personId": 203925,
        "playerName": "Harris",
        "playerNameI": "J. Harris",
        "reboundDefensiveTotal": 1,
        "reboundOffensiveTotal": 0,
        "reboundTotal": 1,
        "subType": "defensive",
        "teamId": 1610612751,
        "teamTricode": "BKN"
      },
      {
        "actionNumber": 40,
        "actionType": "foul",
        "clock": "PT08M46.00S",
        "description": "J. Wiseman shooting personal FOUL (1 PF) (Durant 1 FT)",
        "period": 1,
        "personId": 1630164,
        "playerName": "Wiseman",
        "playerNameI": "J. Wiseman",
        "subType": "personal",
        "teamId": 1610612744,
        "teamTricode": "GSW"
      },
      {
        "actionNumber": 42,
        "actionType": "freethrow",
        "clock": "PT08M46.00S",
        "description": "K. Durant Free Throw 1 of 1 (8 PTS)",
        "period": 1,
        "personId": 201142,
        "playerName": "Durant",
        "playerNameI": "K. Durant",
        "pointsTotal": 8,
        "scoreAway": "6",
        "scoreHome": "13",
        "shotResult": "Made",
        "subType": "1 of 1",
        "teamId": 1610612751,
        "teamTricode": "BKN"
      },
      {
        "actionNumber": 56,
        "actionType": "turnover",
        "clock": "PT07M57.00S",
        "description": "K. Oubre Jr. offensive foul TURNOVER (1 TO)",
        "period": 1,
        "personId": 1626162,
        "playerName": "Oubre Jr.",
        "playerNameI": "K. Oubre Jr.",
        "subType": "offensive foul",
        "teamId": 1610612744,
        "teamTricode": "GSW",
        "turnoverTotal": 1
      },
      {
        "actionNumber": 66,
        "actionType": "substitution",
        "clock": "PT07M20.00S",
        "description": "SUB out: J. Wiseman",
        "period": 1,
        "personId": 1630164,
        "playerName": "Wiseman",
        "playerNameI": "J. Wiseman",
        "subType": "out",
        "teamId": 1610612744,
        "teamTricode": "GSW"
      },
      {
        "actionNumber": 67,
        "actionType": "substitution",
        "clock": "PT07M20.00S",
        "description": "SUB in: K. Looney",
        "period": 1,
        "personId": 1626172,
        "playerName": "Looney",
        "playerNameI": "K. Looney",
        "subType": "in",
        "teamId": 1610612744,
        "teamTricode": "GSW"
      },
      {
        "actionNumber": 729,
        "actionType": "game",
        "clock": "PT00M00.00S",
        "description": "Game End",
        "period": 4,
        "subType": "end"
      }
    ]
  }
}
`
