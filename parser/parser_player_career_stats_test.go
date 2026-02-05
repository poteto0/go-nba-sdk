package parser_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/parser"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

var unexpectedErr *types.GnsError

func TestParsePlayerCareerStatsResponse(t *testing.T) {
	t.Run("should parse normal stats correctly", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource: "playercareerstats",
			ResultSets: []types.ResultSet{
				{
					Name:    "SeasonTotalsRegularSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
				{
					Name:    "CareerTotalsRegularSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
				{
					Name:    "SeasonTotalsPostSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
				{
					Name:    "CareerTotalsPostSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
				{
					Name:    "SeasonTotalsAllStarSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
				{
					Name:    "CareerTotalsAllStarSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
				{
					Name:    "SeasonTotalsShowcaseSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
				{
					Name:    "CareerTotalsShowcaseSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
				{
					Name:    "SeasonTotalsCollegeSeason",
					Headers: constants.PlayerCollegeSeasonStatsRecordHeaders,
					RowSet: [][]any{
						sampleCollegeSeasonStatsRowSet,
					},
				},
				{
					Name:    "CareerTotalsCollegeSeason",
					Headers: constants.PlayerCollegeCareerStatsRecordHeaders,
					RowSet: [][]any{
						sampleCollegeCareerStatsRowSet,
					},
				},
				{
					Name:    "SeasonRankingsRegularSeason",
					Headers: constants.SeasonRankingsRecordHeaders,
					RowSet: [][]any{
						sampleSeasonRankingsRowSet,
					},
				},
				{
					Name:    "SeasonRankingsPostSeason",
					Headers: constants.SeasonRankingsRecordHeaders,
					RowSet: [][]any{
						sampleSeasonRankingsRowSet,
					},
				},
				{
					Name:    "SeasonHighs",
					Headers: constants.PlayerHighsRecordHeaders,
					RowSet: [][]any{
						samplePlayerHighsRowSet,
					},
				},
				{
					Name:    "CareerHighs",
					Headers: constants.PlayerHighsRecordHeaders,
					RowSet: [][]any{
						samplePlayerHighsRowSet,
					},
				},
			},
		}

		// Act
		result, err := parser.ParsePlayerCareerStatsResponse(rawResponse)

		// Assert
		assert.NoError(t, err)

		stats := result.SeasonTotalsRegularSeason[0]
		assert.Equal(t, 203076, stats.PlayerID)
		assert.Equal(t, "2012-13", stats.SeasonID)

		// Optional fields
		v, _ := stats.GP.Take()
		assert.Equal(t, 64, v)

		v, _ = stats.GS.Take()
		assert.Equal(t, 60, v)

		f, _ := stats.Min.Take()
		assert.Equal(t, 28.8, f)

		f, _ = stats.Pts.Take()
		assert.Equal(t, 13.5, f)

		careerStats := result.CareerTotalsRegularSeason[0]
		assert.Equal(t, 203076, careerStats.PlayerID)
		assert.Equal(t, "2012-13", careerStats.SeasonID)

		postStats := result.SeasonTotalsPostSeason[0]
		assert.Equal(t, 203076, postStats.PlayerID)
		assert.Equal(t, "2012-13", postStats.SeasonID)

		allStarStats := result.SeasonTotalsAllStarSeason[0]
		assert.Equal(t, 203076, allStarStats.PlayerID)
		assert.Equal(t, "2012-13", allStarStats.SeasonID)

		careerPostStats := result.CareerTotalsPostSeason[0]
		assert.Equal(t, 203076, careerPostStats.PlayerID)
		assert.Equal(t, "2012-13", careerPostStats.SeasonID)

		careerAllStarStats := result.CareerTotalsAllStarSeason[0]
		assert.Equal(t, 203076, careerAllStarStats.PlayerID)
		assert.Equal(t, "2012-13", careerAllStarStats.SeasonID)

		showcaseStats := result.SeasonTotalsShowcaseSeason[0]
		assert.Equal(t, 203076, showcaseStats.PlayerID)
		assert.Equal(t, "2012-13", showcaseStats.SeasonID)

		careerShowcaseStats := result.CareerTotalsShowcaseSeason[0]
		assert.Equal(t, 203076, careerShowcaseStats.PlayerID)
		assert.Equal(t, "2012-13", careerShowcaseStats.SeasonID)

		collegeStats := result.SeasonTotalsCollegeSeason[0]
		assert.Equal(t, 203076, collegeStats.PlayerID)
		assert.Equal(t, "00", collegeStats.LeagueID)

		// Optional fields
		v2, _ := collegeStats.GP.Take()
		assert.Equal(t, 40, v2)

		f2, _ := collegeStats.Pts.Take()
		assert.Equal(t, 14.2, f2)

		careerCollegeStats := result.CareerTotalsCollegeSeason[0]
		assert.Equal(t, 203076, careerCollegeStats.PlayerID)
		assert.Equal(t, "00", careerCollegeStats.LeagueID)

		rankings := result.SeasonRankingsRegularSeason[0]
		assert.Equal(t, 203076, rankings.PlayerID)
		assert.Equal(t, "2014-15", rankings.SeasonID)

		rankingsPost := result.SeasonRankingsPostSeason[0]
		assert.Equal(t, 203076, rankingsPost.PlayerID)
		assert.Equal(t, "2014-15", rankingsPost.SeasonID)

		highs := result.SeasonHighs[0]
		assert.Equal(t, 203076, highs.PlayerID)
		assert.Equal(t, "0022500404", highs.GameID)

		careerHighs := result.CareerHighs[0]
		assert.Equal(t, 203076, careerHighs.PlayerID)
		assert.Equal(t, "0022500404", careerHighs.GameID)
	})

	t.Run("unexpected resource is error", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource: "unexpected",
			ResultSets: []types.ResultSet{
				{
					Name:    "SeasonTotalsRegularSeason",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
			},
		}

		// Act
		_, err := parser.ParsePlayerCareerStatsResponse(rawResponse)

		// Assert
		assert.ErrorAs(t, err, &unexpectedErr)
		assert.Equal(t, "unexpected resource: unexpected", unexpectedErr.Error())
	})

	t.Run("0 record resultSets is error", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource:   "playercareerstats",
			ResultSets: []types.ResultSet{},
		}

		// Act
		_, err := parser.ParsePlayerCareerStatsResponse(rawResponse)

		// Assert
		assert.ErrorAs(t, err, &unexpectedErr)
		assert.Equal(t, "unexpected no record", unexpectedErr.Error())
	})

	t.Run("unexpected result set name is error", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource: "playercareerstats",
			ResultSets: []types.ResultSet{
				{
					Name:    "unexpected",
					Headers: constants.PlayerCareerStatsRecordHeaders,
					RowSet: [][]any{
						samplePlayerCareerStatsRowSet,
					},
				},
			},
		}

		// Act
		_, err := parser.ParsePlayerCareerStatsResponse(rawResponse)

		// Assert
		assert.ErrorAs(t, err, &unexpectedErr)
		assert.Equal(t, "unexpected result set name: unexpected", unexpectedErr.Error())
	})

	t.Run("should parse rankings with nulls correctly", func(t *testing.T) {
		rawResponse := types.RawResponse{
			Resource: "playercareerstats",
			ResultSets: []types.ResultSet{
				{
					Name: "SeasonRankingsRegularSeason",
					Headers: []string{
						"PLAYER_ID", "SEASON_ID", "LEAGUE_ID", "TEAM_ID", "TEAM_ABBREVIATION",
						"RANK_PG_MIN", "RANK_PG_PTS", "RANK_PG_REB",
					},
					RowSet: [][]any{
						{
							203076, "2012-13", "00", 1610612740, "NOH",
							nil, nil, 85.0, // null, null, 85
						},
					},
				},
			},
		}

		result, err := parser.ParsePlayerCareerStatsResponse(rawResponse)
		assert.NoError(t, err)

		rankings := result.SeasonRankingsRegularSeason[0]

		assert.True(t, rankings.RankPgMin.IsNone())
		assert.True(t, rankings.RankPgPts.IsNone())

		v, _ := rankings.RankPgReb.Take()
		assert.Equal(t, 85, v)
	})

	t.Run("should handle missing stats values as none", func(t *testing.T) {
		// Simulate a case where some stats might be null (unlikely in real API but good for robustness)
		rawResponse := types.RawResponse{
			Resource: "playercareerstats",
			ResultSets: []types.ResultSet{
				{
					Name: "SeasonTotalsRegularSeason",
					Headers: []string{
						"PLAYER_ID", "GP", "MIN",
					},
					RowSet: [][]any{
						{
							203076, nil, nil,
						},
					},
				},
			},
		}

		result, err := parser.ParsePlayerCareerStatsResponse(rawResponse)
		assert.NoError(t, err)

		stats := result.SeasonTotalsRegularSeason[0]
		assert.True(t, stats.GP.IsNone())
		assert.True(t, stats.Min.IsNone())
	})
}
