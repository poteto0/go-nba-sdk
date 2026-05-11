package parser_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/parser"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func TestParseDraftCombineStatsResponse(t *testing.T) {
	t.Run("should parse draft combine stats correctly", func(t *testing.T) {
		headers := []string{
			"SEASON", "PLAYER_ID", "FIRST_NAME", "LAST_NAME", "PLAYER_NAME",
			"POSITION", "HEIGHT_WO_SHOES", "HEIGHT_W_SHOES", "WEIGHT",
			"WINGSPAN", "STANDING_REACH", "BODY_FAT_PCT", "HAND_LENGTH",
			"HAND_WIDTH", "STANDING_VERTICAL_LEAP", "MAX_VERTICAL_LEAP",
			"LANE_AGILITY_TIME", "MODIFIED_LANE_AGILITY_TIME", "THREE_QUARTER_SPRINT",
			"BENCH_PRESS",
		}

		row := []any{
			"2024", 1641700, "Reed", "Sheppard", "Reed Sheppard",
			"G", 73.5, 74.5, 181.6,
			76.5, 96.0, 5.0, 8.5,
			9.0, 31.5, 42.0, 10.5,
			2.8, 3.2, 5,
		}

		rawResponse := types.RawResponse{
			Resource: "draftcombinestats",
			ResultSets: []types.ResultSet{
				{
					Name:    "DraftCombineStats",
					Headers: headers,
					RowSet:  [][]any{row},
				},
			},
		}

		// Act
		result, err := parser.ParseDraftCombineStatsResponse(rawResponse)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(result.CombineStats))

		record := result.CombineStats[0]
		assert.Equal(t, "2024", record.Season)
		assert.Equal(t, 1641700, record.PlayerID)
		assert.Equal(t, "Reed", record.FirstName)
		assert.Equal(t, "Reed Sheppard", record.PlayerName)
		assert.Equal(t, 73.5, *record.HeightWoShoes)
		assert.Equal(t, 5, *record.BenchPress)
	})

	t.Run("unexpected resource is error", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource: "unexpected",
		}

		// Act
		_, err := parser.ParseDraftCombineStatsResponse(rawResponse)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected resource")
	})

	t.Run("0 record resultSets is error", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource:   "draftcombinestats",
			ResultSets: []types.ResultSet{},
		}

		// Act
		_, err := parser.ParseDraftCombineStatsResponse(rawResponse)

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected no record")
	})
}
