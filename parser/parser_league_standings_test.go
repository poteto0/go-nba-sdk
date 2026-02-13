package parser_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/parser"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestParseLeagueStandingsResponse(t *testing.T) {
	t.Run("should parse league standings correctly with all headers", func(t *testing.T) {
		headers := []string{
			"LeagueID", "SeasonID", "TeamID", "TeamCity", "TeamName",
			"Conference", "ConferenceRecord", "PlayoffRank", "ClinchIndicator",
			"Division", "DivisionRecord", "DivisionRank", "WINS", "LOSSES",
			"WinPCT", "LeagueRank", "Record", "HOME", "ROAD", "L10",
			"Last10Home", "Last10Road", "OT", "ThreePTSOrLess", "TenPTSOrMore",
			"LongHomeStreak", "strLongHomeStreak", "LongRoadStreak", "strLongRoadStreak",
			"LongWinStreak", "LongLossStreak", "CurrentHomeStreak", "strCurrentHomeStreak",
			"CurrentRoadStreak", "strCurrentRoadStreak", "CurrentStreak", "strCurrentStreak",
			"ConferenceGamesBack", "DivisionGamesBack", "ClinchedConferenceTitle",
			"ClinchedDivisionTitle", "ClinchedPlayoffBirth", "EliminatedConference",
			"EliminatedDivision", "AheadAtHalf", "BehindAtHalf", "TiedAtHalf",
			"AheadAtThird", "BehindAtThird", "TiedAtThird", "Score100PTS",
			"OppScore100PTS", "OppOver500", "LeadInFGPCT", "LeadInReb",
			"FewerTurnovers", "PointsPG", "OppPointsPG", "DiffPointsPG",
			"vsEast", "vsAtlantic", "vsCentral", "vsSoutheast", "vsWest",
			"vsNorthwest", "vsPacific", "vsSouthwest", "Jan", "Feb", "Mar",
			"Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
			"PreAS", "PostAS",
		}

		row := []any{
			"00", "22025", 1610612760, "Oklahoma City", "Thunder",
			"West", "32-9", 1, "", "Northwest", "8-3", 1, 42, 13,
			0.764, 1, "42-13", "22-5", "20-7", "5-5",
			"6-4", "7-3", "3-0", "2-6", "31-5", 14, "W 14", 8, "W 8",
			16, 2, -1, "L 1", 2, "W 2", 2, "W 2",
			0.0, 0.0, "Title", nil, "Birth", nil, nil,
			"33-7", "9-6", "0-0", "38-3", "4-9", "0-1", "42-12",
			"30-13", "20-11", "35-6", "20-3", "38-8", 120.0, 108.0, 12.0,
			"10-4", "1-1", "3-1", "6-2", "32-9", "8-3", "13-1", "11-5",
			"9-6", "4-2", nil, nil, nil, nil, nil, nil, nil, "6-0", "14-1", "9-4",
			"42-13", nil,
		}

		rawResponse := types.RawResponse{
			Resource: "leaguestandings",
			Parameters: map[string]any{
				"Season": "2025-26",
			},
			ResultSets: []types.ResultSet{
				{
					Name:    "Standings",
					Headers: headers,
					RowSet:  [][]any{row},
				},
			},
		}

		// Act
		result, err := parser.ParseLeagueStandingsResponse(rawResponse)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "2025-26", result.SeasonYear)
		assert.Equal(t, 1, len(result.Standings))
		record := result.Standings[0]

		// Check some key fields
		assert.Equal(t, "00", record.LeagueID)
		assert.Equal(t, "Thunder", record.TeamName)
		assert.Equal(t, 42, record.Wins)
		assert.Equal(t, 120.0, record.PointsPG)

		// Check optional fields
		assert.NotNil(t, record.ClinchedConferenceTitle)
		assert.Equal(t, "Title", *record.ClinchedConferenceTitle)

		assert.Nil(t, record.ClinchedDivisionTitle)

		assert.NotNil(t, record.ClinchedPlayoffBirth)
		assert.Equal(t, "Birth", *record.ClinchedPlayoffBirth)

		// Check monthly/break fields
		assert.NotNil(t, record.Jan)
		assert.Equal(t, "9-6", *record.Jan)

		assert.Nil(t, record.Mar)

		assert.NotNil(t, record.PreAS)
		assert.Equal(t, "42-13", *record.PreAS)

		assert.Nil(t, record.PostAS)
	})

	t.Run("unexpected resource is error", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource: "unexpected",
		}

		// Act
		_, err := parser.ParseLeagueStandingsResponse(rawResponse)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected resource")
	})

	t.Run("0 record resultSets is error", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource:   "leaguestandings",
			ResultSets: []types.ResultSet{},
		}

		// Act
		_, err := parser.ParseLeagueStandingsResponse(rawResponse)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected no record")
	})
}
