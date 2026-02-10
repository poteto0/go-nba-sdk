package samples

var SampleScoreBoardResponse = `
{
  "meta": {
    "version": 1,
    "request": "https://nba-prod-us-east-1-mediaops-stats.s3.amazonaws.com/NBA/liveData/scoreboard/todaysScoreboard_00.json",
    "time": "2026-02-06 06:27:33.2733",
    "code": 200
  },
  "scoreboard": {
    "gameDate": "2026-02-05",
    "leagueId": "00",
    "leagueName": "National Basketball Association",
    "games": [
      {
        "gameId": "0022500733",
        "gameCode": "20260205/WASDET",
        "gameStatus": 3,
        "gameStatusText": "Final",
        "period": 4,
        "gameClock": "",
        "gameTimeUTC": "2026-02-06T00:00:00Z",
        "gameEt": "2026-02-05T19:00:00Z",
        "regulationPeriods": 4,
        "ifNecessary": false,
        "seriesGameNumber": "",
        "gameLabel": "",
        "gameSubLabel": "",
        "seriesText": "",
        "seriesConference": "",
        "poRoundDesc": "",
        "gameSubtype": "",
        "isNeutral": false,
        "homeTeam": {
          "teamId": 1610612765,
          "teamName": "Pistons",
          "teamCity": "Detroit",
          "teamTricode": "DET",
          "wins": 37,
          "losses": 13,
          "score": 117,
          "seed": null,
          "inBonus": null,
          "timeoutsRemaining": 0,
          "periods": [
            {
              "period": 1,
              "periodType": "REGULAR",
              "score": 21
            },
            {
              "period": 2,
              "periodType": "REGULAR",
              "score": 31
            },
            {
              "period": 3,
              "periodType": "REGULAR",
              "score": 32
            },
            {
              "period": 4,
              "periodType": "REGULAR",
              "score": 33
            }
          ]
        },
        "awayTeam": {
          "teamId": 1610612764,
          "teamName": "Wizards",
          "teamCity": "Washington",
          "teamTricode": "WAS",
          "wins": 14,
          "losses": 36,
          "score": 126,
          "seed": null,
          "inBonus": null,
          "timeoutsRemaining": 1,
          "periods": [
            {
              "period": 1,
              "periodType": "REGULAR",
              "score": 34
            },
            {
              "period": 2,
              "periodType": "REGULAR",
              "score": 22
            },
            {
              "period": 3,
              "periodType": "REGULAR",
              "score": 39
            },
            {
              "period": 4,
              "periodType": "REGULAR",
              "score": 31
            }
          ]
        },
        "gameLeaders": {
          "homeLeaders": {
            "personId": 1630595,
            "name": "Cade Cunningham",
            "jerseyNum": "2",
            "position": "G",
            "teamTricode": "DET",
            "playerSlug": null,
            "points": 30,
            "rebounds": 8,
            "assists": 8
          },
          "awayLeaders": {
            "personId": 1642860,
            "name": "Will Riley",
            "jerseyNum": "27",
            "position": "F",
            "teamTricode": "WAS",
            "playerSlug": null,
            "points": 20,
            "rebounds": 6,
            "assists": 5
          }
        },
        "pbOdds": {
          "team": null,
          "odds": 0.0,
          "suspended": 0
        }
      }
    ]
  }
}
`
