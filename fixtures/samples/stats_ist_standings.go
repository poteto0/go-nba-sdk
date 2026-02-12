package samples

var SampleIstStandingsResponse = `{
    "resource": "iststandings",
    "parameters": {
        "LeagueID": "00",
        "Season": "2025-26"
    },
    "resultSets": [
        {
            "name": "Standings",
            "headers": [
                "LeagueId",
                "SeasonID",
                "TeamID",
                "TeamCity",
                "TeamName",
                "TeamAbbreviation",
                "Conference",
                "Group",
                "PlayoffRank",
                "W",
                "L",
                "W_PCT",
                "HOME",
                "ROAD",
                "OT",
                "LAST10",
                "STREAK",
                "ConferenceRecord",
                "DivisionRecord",
                "GroupRecord",
                "ClinchedPlayoff",
                "ClinchedPlayIn",
                "ClinchedConference",
                "ClinchedDivision",
                "ClinchedGroup",
                "EliminatedPlayoff",
                "EliminatedPlayIn",
                "EliminatedConference",
                "EliminatedDivision",
                "EliminatedGroup"
            ],
            "rowSet": [
                [
                    "00",
                    "22023",
                    1610612747,
                    "Los Angeles",
                    "Lakers",
                    "LAL",
                    "West",
                    "West Group A",
                    1,
                    4,
                    0,
                    1.0,
                    "2-0",
                    "2-0",
                    "0-0",
                    "4-0",
                    "W 4",
                    "4-0",
                    "4-0",
                    "4-0",
                    0,
                    0,
                    0,
                    0,
                    1,
                    0,
                    0,
                    0,
                    0,
                    0
                ]
            ]
        }
    ]
}`

var SampleInvalidIstStandingsResponse = `{
    "resource": "iststandings",
    "parameters": {
        "LeagueID": "00",
        "Season": "2025-26"
    },
    "resultSets": [
        "hello",
    ]
}`
