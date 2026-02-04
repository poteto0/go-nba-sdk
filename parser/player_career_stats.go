package parser

import (
	"errors"
	"fmt"

	"github.com/poteto0/go-nba-sdk/types"
)

func ParsePlayerCareerStatsResponse(rawResponse types.RawResponse) (types.PlayerCareerStatsResponseContent, error) {
	if rawResponse.Resource != "playercareerstats" {
		return types.PlayerCareerStatsResponseContent{}, errors.New("unexpected")
	}

	if len(rawResponse.ResultSets) == 0 {
		return types.PlayerCareerStatsResponseContent{}, errors.New("no result sets")
	}

	content := types.PlayerCareerStatsResponseContent{}
	for _, resultSet := range rawResponse.ResultSets {
		switch resultSet.Name {
		case "SeasonTotalsRegularSeason":
			content.SeasonTotalsRegularSeason = parsePlayerCareerStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "CareerTotalsRegularSeason":
			content.CareerTotalsRegularSeason = parsePlayerCareerStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonTotalsPostSeason":
			content.SeasonTotalsPostSeason = parsePlayerCareerStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "CareerTotalsPostSeason":
			content.CareerTotalsPostSeason = parsePlayerCareerStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonTotalsAllStarSeason":
			content.SeasonTotalsAllStarSeason = parsePlayerCareerStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "CareerTotalsAllStarSeason":
			content.CareerTotalsAllStarSeason = parsePlayerCareerStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonTotalsShowcaseSeason":
			content.SeasonTotalsShowcaseSeason = parsePlayerCareerStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "CareerTotalsShowcaseSeason":
			content.CareerTotalsShowcaseSeason = parsePlayerCareerStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonTotalsCollegeSeason":
			content.SeasonTotalsCollegeSeason = parsePlayerCollegeStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "CareerTotalsCollegeSeason":
			content.CareerTotalsCollegeSeason = parsePlayerCollegeStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonRankingsRegularSeason":
			content.SeasonRankingsRegularSeason = parsePlayerCareerRankingsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonRankingsPostSeason":
			content.SeasonRankingsPostSeason = parsePlayerCareerRankingsRecords(resultSet.RowSet, resultSet.Headers)

		//case "SeasonHighs":
		// TODO: ほんとはエラーしたいが、パースできていないエラーが多い
		default:
			continue
		}
	}

	return content, nil
}

// analyze headers to find the index of each field
func parsePlayerCareerStatsRecords(
	rawSet [][]any, headers []string,
) []types.PlayerCareerStatsRecord {
	records := make([]types.PlayerCareerStatsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.PlayerCareerStatsRecord
		for j, header := range headers {
			switch header {
			case "PLAYER_ID":
				record.PlayerID = int(row[j].(float64))
			case "SEASON_ID":
				record.SeasonID = row[j].(string)
			case "LEAGUE_ID":
				record.LeagueID = row[j].(string)
			case "TEAM_ID":
				record.TeamID = int(row[j].(float64))
			case "TEAM_ABBREVIATION":
				record.TeamAbbeviation = row[j].(string)
			case "PLAYER_AGE":
				record.PlayerAge = int(row[j].(float64))
			case "GP":
				record.GP = int(row[j].(float64))
			case "GS":
				record.GS = int(row[j].(float64))
			case "MIN":
				record.MIN = row[j].(float64)
			case "PTS":
				record.PTS = row[j].(float64)
			case "FGM":
				record.FGM = row[j].(float64)
			case "FGA":
				record.FGA = row[j].(float64)
			case "FG_PCT":
				record.FG_PCT = row[j].(float64)
			case "FG3M":
				record.FG3M = row[j].(float64)
			case "FG3A":
				record.FG3A = row[j].(float64)
			case "FG3_PCT":
				record.FG3_PCT = row[j].(float64)
			case "FTM":
				record.FTM = row[j].(float64)
			case "FTA":
				record.FTA = row[j].(float64)
			case "FT_PCT":
				record.FT_PCT = row[j].(float64)
			case "OREB":
				record.OREB = row[j].(float64)
			case "DREB":
				record.DREB = row[j].(float64)
			case "REB":
				record.REB = row[j].(float64)
			case "AST":
				record.AST = row[j].(float64)
			case "STL":
				record.STL = row[j].(float64)
			case "BLK":
				record.BLK = row[j].(float64)
			case "TOV":
				record.TOV = row[j].(float64)
			case "PF":
				record.PF = row[j].(float64)
			}
		}
		records[i] = record
	}

	return records
}

func parsePlayerCollegeStatsRecords(
	rawSet [][]any, headers []string,
) []types.PlayerCollegeStatsRecord {
	records := make([]types.PlayerCollegeStatsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.PlayerCollegeStatsRecord
		for j, header := range headers {
			switch header {
			case "PLAYER_ID":
				record.PlayerID = int(row[j].(float64))
			case "SEASON_ID":
				record.SeasonID = row[j].(string)
			case "LEAGUE_ID":
				record.LeagueID = row[j].(string)
			case "ORGANIZATION_ID":
				record.ORGANIZATION_ID = row[j].(float64)
			case "SCHOOL_NAME":
				record.SCHOOL_NAME = row[j].(string)
			case "GP":
				record.GP = int(row[j].(float64))
			case "GS":
				record.GS = int(row[j].(float64))
			case "MIN":
				record.MIN = row[j].(float64)
			case "PTS":
				record.PTS = row[j].(float64)
			case "FGM":
				record.FGM = row[j].(float64)
			case "FGA":
				record.FGA = row[j].(float64)
			case "FG_PCT":
				record.FG_PCT = row[j].(float64)
			case "FG3M":
				record.FG3M = row[j].(float64)
			case "FG3A":
				record.FG3A = row[j].(float64)
			case "FG3_PCT":
				record.FG3_PCT = row[j].(float64)
			case "FTM":
				record.FTM = row[j].(float64)
			case "FTA":
				record.FTA = row[j].(float64)
			case "FT_PCT":
				record.FT_PCT = row[j].(float64)
			case "OREB":
				record.OREB = row[j].(float64)
			case "DREB":
				record.DREB = row[j].(float64)
			case "REB":
				record.REB = row[j].(float64)
			case "AST":
				record.AST = row[j].(float64)
			case "STL":
				record.STL = row[j].(float64)
			case "BLK":
				record.BLK = row[j].(float64)
			case "TOV":
				record.TOV = row[j].(float64)
			case "PF":
				record.PF = row[j].(float64)
			}
		}
		records[i] = record
	}

	return records
}

func parsePlayerCareerRankingsRecords(
	rawSet [][]any, headers []string,
) []types.PlayerCareerRankingsRecord {
	records := make([]types.PlayerCareerRankingsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.PlayerCareerRankingsRecord
		for j, header := range headers {
			switch header {
			case "PLAYER_ID":
				record.PlayerID = int(row[j].(float64))
			case "SEASON_ID":
				record.SeasonID = row[j].(string)
			case "LEAGUE_ID":
				record.LeagueID = row[j].(string)
			case "TEAM_ID":
				record.TeamID = int(row[j].(float64))
			case "TEAM_ABBREVIATION":
				record.TeamAbbeviation = row[j].(string)
			case "RANK_PG_MIN":
				// nil -> 0
				if row[j] != nil {
					record.RankPgMin = int(row[j].(float64))
				}
			case "RANK_PG_PTS":
				if row[j] != nil {
					fmt.Println(int(row[j].(float64)))
					record.RankPgPts = int(row[j].(float64))
				}
			case "RANK_PG_OREB":
				if row[j] != nil {
					record.RankPgOReb = int(row[j].(float64))
				}
			case "RANK_PG_DREB":
				if row[j] != nil {
					record.RankPgDReb = int(row[j].(float64))
				}
			case "RANK_PG_REB":
				if row[j] != nil {
					record.RankPgReb = int(row[j].(float64))
				}
			case "RANK_PG_AST":
				if row[j] != nil {
					record.RankPgAst = int(row[j].(float64))
				}
			case "RANK_PG_STL":
				if row[j] != nil {
					record.RankPgStl = int(row[j].(float64))
				}
			case "RANK_PG_BLK":
				if row[j] != nil {
					record.RankPgBlk = int(row[j].(float64))
				}
			case "RANK_PG_TOV":
				if row[j] != nil {
					record.RankPgTov = int(row[j].(float64))
				}
			case "RANK_PG_FGM":
				if row[j] != nil {
					record.RankPgFgm = int(row[j].(float64))
				}
			case "RANK_PG_FGA":
				if row[j] != nil {
					record.RankPgFga = int(row[j].(float64))
				}
			case "RANK_PG_FG_PCT":
				if row[j] != nil {
					record.RankPgFgPct = int(row[j].(float64))
				}
			case "RANK_PG_FTM":
				if row[j] != nil {
					record.RankPgFtm = int(row[j].(float64))
				}
			case "RANK_PG_FTA":
				if row[j] != nil {
					record.RankPgFta = int(row[j].(float64))
				}
			case "RANK_PG_FT_PCT":
				if row[j] != nil {
					record.RankPgFtPct = int(row[j].(float64))
				}
			case "RANK_PG_EFF":
				if row[j] != nil {
					record.RankPgEff = int(row[j].(float64))
				}
			case "PLAYER_AGE":
				// ignore
			case "GP":
				// ignore
			case "GS":
				// ignore
			}
		}
		records[i] = record
	}

	return records
}
