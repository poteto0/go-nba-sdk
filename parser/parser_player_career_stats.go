package parser

import (
	"github.com/poteto0/go-nba-sdk/types"
)

func ParsePlayerCareerStatsResponse(rawResponse types.RawResponse) (types.PlayerCareerStatsResponseContent, error) {
	if rawResponse.Resource != "playercareerstats" {
		return types.PlayerCareerStatsResponseContent{}, types.NewGnsError(
			"unexpected resource: %s", rawResponse.Resource,
		)
	}

	if len(rawResponse.ResultSets) == 0 {
		return types.PlayerCareerStatsResponseContent{}, types.NewGnsError(
			"unexpected no record",
		)
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
			content.SeasonTotalsCollegeSeason = parsePlayerSeasonCollegeStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "CareerTotalsCollegeSeason":
			content.CareerTotalsCollegeSeason = parsePlayerCareerCollegeStatsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonRankingsRegularSeason":
			content.SeasonRankingsRegularSeason = parsePlayerCareerRankingsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonRankingsPostSeason":
			content.SeasonRankingsPostSeason = parsePlayerCareerRankingsRecords(resultSet.RowSet, resultSet.Headers)

		case "SeasonHighs":
			content.SeasonHighs = parsePlayerHighsRecords(resultSet.RowSet, resultSet.Headers)

		case "CareerHighs":
			content.CareerHighs = parsePlayerHighsRecords(resultSet.RowSet, resultSet.Headers)

		default:
			return types.PlayerCareerStatsResponseContent{}, types.NewGnsError(
				"unexpected result set name: %s", resultSet.Name,
			)
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
				record.PlayerID = toInt(row[j])
			case "SEASON_ID":
				record.SeasonID = row[j].(string)
			case "LEAGUE_ID":
				record.LeagueID = row[j].(string)
			case "TEAM_ID":
				record.TeamID = toInt(row[j])
			case "TEAM_ABBREVIATION":
				record.TeamAbbeviation = row[j].(string)
			case "PLAYER_AGE":
				record.PlayerAge = toInt(row[j])
			case "GP":
				record.GP = toPtrInt(row[j])
			case "GS":
				record.GS = toPtrInt(row[j])
			case "MIN":
				record.Min = toPtrFloat(row[j])
			case "PTS":
				record.Pts = toPtrFloat(row[j])
			case "FGM":
				record.FgM = toPtrFloat(row[j])
			case "FGA":
				record.FgA = toPtrFloat(row[j])
			case "FG_PCT":
				record.FgPct = toPtrFloat(row[j])
			case "FG3M":
				record.Fg3M = toPtrFloat(row[j])
			case "FG3A":
				record.Fg3A = toPtrFloat(row[j])
			case "FG3_PCT":
				record.Fg3Pct = toPtrFloat(row[j])
			case "FTM":
				record.FtM = toPtrFloat(row[j])
			case "FTA":
				record.FtA = toPtrFloat(row[j])
			case "FT_PCT":
				record.FtPct = toPtrFloat(row[j])
			case "OREB":
				record.OReb = toPtrFloat(row[j])
			case "DREB":
				record.DReb = toPtrFloat(row[j])
			case "REB":
				record.Reb = toPtrFloat(row[j])
			case "AST":
				record.Ast = toPtrFloat(row[j])
			case "STL":
				record.Stl = toPtrFloat(row[j])
			case "BLK":
				record.Blk = toPtrFloat(row[j])
			case "TOV":
				record.Tov = toPtrFloat(row[j])
			case "PF":
				record.PF = toPtrFloat(row[j])
			}
		}
		records[i] = record
	}

	return records
}

func parsePlayerSeasonCollegeStatsRecords(
	rawSet [][]any, headers []string,
) []types.PlayerCollegeSeasonStatsRecord {
	records := make([]types.PlayerCollegeSeasonStatsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.PlayerCollegeSeasonStatsRecord
		for j, header := range headers {
			switch header {
			case "PLAYER_ID":
				record.PlayerID = toInt(row[j])
			case "SEASON_ID":
				record.SeasonID = row[j].(string)
			case "LEAGUE_ID":
				record.LeagueID = row[j].(string)
			case "ORGANIZATION_ID":
				record.OrganizationID = toInt(row[j])
			case "SCHOOL_NAME":
				record.SchoolName = row[j].(string)
			case "GP":
				record.GP = toPtrInt(row[j])
			case "GS":
				record.GS = toPtrInt(row[j])
			case "MIN":
				record.Min = toPtrFloat(row[j])
			case "PTS":
				record.Pts = toPtrFloat(row[j])
			case "FGM":
				record.FgM = toPtrFloat(row[j])
			case "FGA":
				record.FgA = toPtrFloat(row[j])
			case "FG_PCT":
				record.FgPct = toPtrFloat(row[j])
			case "FG3M":
				record.Fg3M = toPtrFloat(row[j])
			case "FG3A":
				record.Fg3A = toPtrFloat(row[j])
			case "FG3_PCT":
				record.Fg3Pct = toPtrFloat(row[j])
			case "FTM":
				record.FtM = toPtrFloat(row[j])
			case "FTA":
				record.FtA = toPtrFloat(row[j])
			case "FT_PCT":
				record.FtPct = toPtrFloat(row[j])
			case "OREB":
				record.OReb = toPtrFloat(row[j])
			case "DREB":
				record.DReb = toPtrFloat(row[j])
			case "REB":
				record.Reb = toPtrFloat(row[j])
			case "AST":
				record.Ast = toPtrFloat(row[j])
			case "STL":
				record.Stl = toPtrFloat(row[j])
			case "BLK":
				record.Blk = toPtrFloat(row[j])
			case "TOV":
				record.Tov = toPtrFloat(row[j])
			case "PF":
				record.PF = toPtrFloat(row[j])
			}
		}
		records[i] = record
	}

	return records
}

func parsePlayerCareerCollegeStatsRecords(
	rawSet [][]any, headers []string,
) []types.PlayerCollegeSeasonStatsRecord {
	records := make([]types.PlayerCollegeSeasonStatsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.PlayerCollegeSeasonStatsRecord
		for j, header := range headers {
			switch header {
			case "PLAYER_ID":
				record.PlayerID = toInt(row[j])
			case "LEAGUE_ID":
				record.LeagueID = row[j].(string)
			case "ORGANIZATION_ID":
				record.OrganizationID = toInt(row[j])
			case "GP":
				record.GP = toPtrInt(row[j])
			case "GS":
				record.GS = toPtrInt(row[j])
			case "MIN":
				record.Min = toPtrFloat(row[j])
			case "PTS":
				record.Pts = toPtrFloat(row[j])
			case "FGM":
				record.FgM = toPtrFloat(row[j])
			case "FGA":
				record.FgA = toPtrFloat(row[j])
			case "FG_PCT":
				record.FgPct = toPtrFloat(row[j])
			case "FG3M":
				record.Fg3M = toPtrFloat(row[j])
			case "FG3A":
				record.Fg3A = toPtrFloat(row[j])
			case "FG3_PCT":
				record.Fg3Pct = toPtrFloat(row[j])
			case "FTM":
				record.FtM = toPtrFloat(row[j])
			case "FTA":
				record.FtA = toPtrFloat(row[j])
			case "FT_PCT":
				record.FtPct = toPtrFloat(row[j])
			case "OREB":
				record.OReb = toPtrFloat(row[j])
			case "DREB":
				record.DReb = toPtrFloat(row[j])
			case "REB":
				record.Reb = toPtrFloat(row[j])
			case "AST":
				record.Ast = toPtrFloat(row[j])
			case "STL":
				record.Stl = toPtrFloat(row[j])
			case "BLK":
				record.Blk = toPtrFloat(row[j])
			case "TOV":
				record.Tov = toPtrFloat(row[j])
			case "PF":
				record.PF = toPtrFloat(row[j])
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
				record.PlayerID = toInt(row[j])
			case "SEASON_ID":
				record.SeasonID = row[j].(string)
			case "LEAGUE_ID":
				record.LeagueID = row[j].(string)
			case "TEAM_ID":
				record.TeamID = toInt(row[j])
			case "TEAM_ABBREVIATION":
				record.TeamAbbeviation = row[j].(string)
			case "RANK_PG_MIN":
				record.RankPgMin = toPtrInt(row[j])
			case "RANK_PG_PTS":
				record.RankPgPts = toPtrInt(row[j])
			case "RANK_PG_OREB":
				record.RankPgOReb = toPtrInt(row[j])
			case "RANK_PG_DREB":
				record.RankPgDReb = toPtrInt(row[j])
			case "RANK_PG_REB":
				record.RankPgReb = toPtrInt(row[j])
			case "RANK_PG_AST":
				record.RankPgAst = toPtrInt(row[j])
			case "RANK_PG_STL":
				record.RankPgStl = toPtrInt(row[j])
			case "RANK_PG_BLK":
				record.RankPgBlk = toPtrInt(row[j])
			case "RANK_PG_TOV":
				record.RankPgTov = toPtrInt(row[j])
			case "RANK_PG_FGM":
				record.RankPgFgm = toPtrInt(row[j])
			case "RANK_PG_FGA":
				record.RankPgFga = toPtrInt(row[j])
			case "RANK_FG_PCT":
				record.RankPgFgPct = toPtrInt(row[j])
			case "RANK_PG_FG3M":
				record.RankPgFg3m = toPtrInt(row[j])
			case "RANK_PG_FG3A":
				record.RankPgFg3a = toPtrInt(row[j])
			case "RANK_FG3_PCT":
				record.RankPgFg3Pct = toPtrInt(row[j])
			case "RANK_PG_FTM":
				record.RankPgFtm = toPtrInt(row[j])
			case "RANK_PG_FTA":
				record.RankPgFta = toPtrInt(row[j])
			case "RANK_FT_PCT":
				record.RankPgFtPct = toPtrInt(row[j])
			case "RANK_PG_EFF":
				record.RankPgEff = toPtrInt(row[j])
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

func parsePlayerHighsRecords(
	rawSet [][]any, headers []string,
) []types.PlayerHighsRecord {
	records := make([]types.PlayerHighsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.PlayerHighsRecord
		for j, header := range headers {
			switch header {
			case "PLAYER_ID":
				record.PlayerID = toInt(row[j])
			case "GAME_ID":
				record.GameID = row[j].(string)
			case "GAME_DATE":
				record.GameDate = row[j].(string)
			case "VS_TEAM_ID":
				record.VsTeamID = toInt(row[j])
			case "VS_TEAM_ABBREVIATION":
				record.VsTeamAbbreviation = row[j].(string)
			case "STAT":
				record.Stat = row[j].(string)
			case "STAT_VALUE":
				record.StatValue = toFloat(row[j])
			case "STAT_ORDER":
				record.StatOrder = toInt(row[j])
			case "DATE_EST":
				record.DateEst = row[j].(string)
			}
		}
		records[i] = record
	}

	return records
}
