package parser

import (
	"github.com/poteto0/go-nba-sdk/types"
)

func ParseIstStandingsResponse(rawResponse types.RawResponse) (types.IstStandingsResponseContent, error) {
	if rawResponse.Resource != "iststandings" {
		return types.IstStandingsResponseContent{}, types.NewGnsError(
			"unexpected resource: %s", rawResponse.Resource,
		)
	}

	if len(rawResponse.ResultSets) == 0 {
		return types.IstStandingsResponseContent{}, types.NewGnsError(
			"unexpected no record",
		)
	}

	content := types.IstStandingsResponseContent{}
	for _, resultSet := range rawResponse.ResultSets {
		if resultSet.Name == "Standings" {
			content.Standings = parseIstStandingsRecords(resultSet.RowSet, resultSet.Headers)
		}
	}

	return content, nil
}

func parseIstStandingsRecords(
	rawSet [][]any, headers []string,
) []types.IstStandingsRecord {
	records := make([]types.IstStandingsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.IstStandingsRecord
		for j, header := range headers {
			switch header {
			case "LeagueId":
				record.LeagueId = row[j].(string)
			case "SeasonID":
				record.SeasonID = row[j].(string)
			case "TeamID":
				record.TeamID = toInt(row[j])
			case "TeamCity":
				record.TeamCity = row[j].(string)
			case "TeamName":
				record.TeamName = row[j].(string)
			case "TeamAbbreviation":
				record.TeamAbbreviation = row[j].(string)
			case "Conference":
				record.Conference = row[j].(string)
			case "Group":
				record.Group = row[j].(string)
			case "PlayoffRank":
				record.PlayoffRank = toInt(row[j])
			case "W":
				record.W = toInt(row[j])
			case "L":
				record.L = toInt(row[j])
			case "W_PCT":
				record.WPct = toFloat(row[j])
			case "HOME":
				record.Home = row[j].(string)
			case "ROAD":
				record.Road = row[j].(string)
			case "OT":
				record.OT = row[j].(string)
			case "LAST10":
				record.Last10 = row[j].(string)
			case "STREAK":
				record.Streak = row[j].(string)
			case "ConferenceRecord":
				record.ConferenceRecord = row[j].(string)
			case "DivisionRecord":
				record.DivisionRecord = row[j].(string)
			case "GroupRecord":
				record.GroupRecord = row[j].(string)
			case "ClinchedPlayoff":
				record.ClinchedPlayoff = toInt(row[j])
			case "ClinchedPlayIn":
				record.ClinchedPlayIn = toInt(row[j])
			case "ClinchedConference":
				record.ClinchedConference = toInt(row[j])
			case "ClinchedDivision":
				record.ClinchedDivision = toInt(row[j])
			case "ClinchedGroup":
				record.ClinchedGroup = toInt(row[j])
			case "EliminatedPlayoff":
				record.EliminatedPlayoff = toInt(row[j])
			case "EliminatedPlayIn":
				record.EliminatedPlayIn = toInt(row[j])
			case "EliminatedConference":
				record.EliminatedConference = toInt(row[j])
			case "EliminatedDivision":
				record.EliminatedDivision = toInt(row[j])
			case "EliminatedGroup":
				record.EliminatedGroup = toInt(row[j])
			}
		}
		records[i] = record
	}

	return records
}
