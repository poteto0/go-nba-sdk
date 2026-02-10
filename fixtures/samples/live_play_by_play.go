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
      }
    ]
  }
}
`
