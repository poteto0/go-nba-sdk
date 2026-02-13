package samples

var SampleScheduleLeagueV2Response = `{
	"meta": {
		"version": 1,
		"request": "http://nba.cloud/league/00/2024-25/scheduleleaguev2?Format=json",
		"time": "2025-10-27T10:36:14.3614Z"
	},
	"leagueSchedule": {
		"seasonYear": "2024-25",
		"leagueId": "00",
		"gameDates": [
			{
				"gameDate": "10/04/2024 00:00:00",
				"games": [
					{
						"gameId": "0012400001",
						"gameCode": "20241004/BOSDEN",
						"gameStatus": 3,
						"gameStatusText": "Final",
						"gameSequence": 1,
						"gameDateEst": "2024-10-04T00:00:00Z",
						"gameTimeEst": "1900-01-01T12:00:00Z",
						"gameDateTimeEst": "2024-10-04T12:00:00Z",
						"gameDateUTC": "2024-10-04T04:00:00Z",
						"gameTimeUTC": "1900-01-01T16:00:00Z",
						"gameDateTimeUTC": "2024-10-04T16:00:00Z",
						"awayTeamTime": "2024-10-04T12:00:00Z",
						"homeTeamTime": "2024-10-04T10:00:00Z",
						"day": "Fri",
						"monthNum": 10,
						"weekNumber": 0,
						"weekName": "",
						"ifNecessary": "false",
						"seriesGameNumber": "",
						"gameLabel": "Preseason",
						"gameSubLabel": "NBA Abu Dhabi Game",
						"seriesText": "Neutral Site",
						"arenaName": "Etihad Arena",
						"arenaState": "",
						"arenaCity": "Abu Dhabi",
						"postponedStatus": "N",
						"branchLink": "",
						"gameSubtype": "Global Games",
						"isNeutral": true,
						"broadcasters": {
							"nationalBroadcasters": [
								{
									"broadcasterScope": "natl",
									"broadcasterMedia": "tv",
									"broadcasterId": 7,
									"broadcasterDisplay": "NBA TV",
									"broadcasterAbbreviation": "NBA TV",
									"broadcasterDescription": "",
									"tapeDelayComments": "",
									"broadcasterVideoLink": "",
									"broadcasterTeamId": -1,
									"broadcasterRanking": null,
									"localizationRegion": ""
								}
							],
							"nationalRadioBroadcasters": [],
							"nationalOttBroadcasters": [],
							"homeTvBroadcasters": [],
							"homeRadioBroadcasters": [],
							"homeOttBroadcasters": [],
							"awayTvBroadcasters": [],
							"awayRadioBroadcasters": [],
							"awayOttBroadcasters": []
						},
						"homeTeam": {
							"teamId": 1610612743,
							"teamName": "Nuggets",
							"teamCity": "Denver",
							"teamTricode": "DEN",
							"teamSlug": "nuggets",
							"wins": 0,
							"losses": 1,
							"score": 103,
							"seed": 0
						},
						"awayTeam": {
							"teamId": 1610612738,
							"teamName": "Celtics",
							"teamCity": "Boston",
							"teamTricode": "BOS",
							"teamSlug": "celtics",
							"wins": 1,
							"losses": 0,
							"score": 107,
							"seed": 0
						},
						"pointsLeaders": [
							{
								"personId": 1630202,
								"firstName": "Payton",
								"lastName": "Pritchard",
								"teamId": 1610612738,
								"teamCity": "Boston",
								"teamName": "Celtics",
								"teamTricode": "BOS",
								"points": 21.0
							}
						]
					}
				]
			}
		]
	}
}`
