package parser_test

import (
	"encoding/json"
	"testing"

	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/poteto0/go-nba-sdk/parser"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_ParseIstStandingsResponse(t *testing.T) {
	t.Run("can parse ist standings response", func(t *testing.T) {
		// Arrange
		var rawResponse types.RawResponse
		err := json.Unmarshal([]byte(samples.SampleIstStandingsResponse), &rawResponse)
		assert.NoError(t, err)

		// Act
		content, err := parser.ParseIstStandingsResponse(rawResponse)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(content.Standings))

		record := content.Standings[0]
		assert.Equal(t, "00", record.LeagueId)
		assert.Equal(t, "22023", record.SeasonID)
		assert.Equal(t, 1610612747, record.TeamID)
		assert.Equal(t, "Los Angeles", record.TeamCity)
		assert.Equal(t, "Lakers", record.TeamName)
		assert.Equal(t, "LAL", record.TeamAbbreviation)
		assert.Equal(t, "West", record.Conference)
		assert.Equal(t, "West Group A", record.Group)
		assert.Equal(t, 1, record.PlayoffRank)
		assert.Equal(t, 4, record.W)
		assert.Equal(t, 0, record.L)
		assert.Equal(t, 1.0, record.WPct)
		assert.Equal(t, "2-0", record.Home)
		assert.Equal(t, "2-0", record.Road)
		assert.Equal(t, "0-0", record.OT)
		assert.Equal(t, "4-0", record.Last10)
		assert.Equal(t, "W 4", record.Streak)
		assert.Equal(t, "4-0", record.ConferenceRecord)
		assert.Equal(t, "4-0", record.DivisionRecord)
		assert.Equal(t, "4-0", record.GroupRecord)
		assert.Equal(t, 1, record.ClinchedGroup)
	})

	t.Run("error on unexpected resource", func(t *testing.T) {
		// Arrange
		rawResponse := types.RawResponse{
			Resource: "unexpected",
		}

		// Act
		_, err := parser.ParseIstStandingsResponse(rawResponse)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected resource")
	})

	t.Run("error on no record", func(t *testing.T) {
		// Arrange
		rawResponse := types.RawResponse{
			Resource: "iststandings",
		}

		// Act
		_, err := parser.ParseIstStandingsResponse(rawResponse)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected no record")
	})
}
