package parser

import (
	"github.com/poteto0/go-nba-sdk/types"
)

func ParseDraftCombineStatsResponse(rawResponse types.RawResponse) (types.DraftCombineStatsResponseContent, error) {
	if rawResponse.Resource != "draftcombinestats" {
		return types.DraftCombineStatsResponseContent{}, types.NewGnsError(
			"unexpected resource: %s", rawResponse.Resource,
		)
	}

	if len(rawResponse.ResultSets) == 0 {
		return types.DraftCombineStatsResponseContent{}, types.NewGnsError(
			"unexpected no record",
		)
	}

	content := types.DraftCombineStatsResponseContent{}

	for _, resultSet := range rawResponse.ResultSets {
		if resultSet.Name == "DraftCombineStats" {
			content.CombineStats = parseDraftCombineStatsRecords(resultSet.RowSet, resultSet.Headers)
		}
	}

	return content, nil
}

func parseDraftCombineStatsRecords(rawSet [][]any, headers []string) []types.DraftCombineStatsRecord {
	records := make([]types.DraftCombineStatsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.DraftCombineStatsRecord
		for j, header := range headers {
			val := row[j]
			if val == nil {
				continue
			}
			switch header {
			case "SEASON":
				record.Season = val.(string)
			case "PLAYER_ID":
				record.PlayerID = toInt(val)
			case "FIRST_NAME":
				record.FirstName = val.(string)
			case "LAST_NAME":
				record.LastName = val.(string)
			case "PLAYER_NAME":
				record.PlayerName = val.(string)
			case "POSITION":
				record.Position = val.(string)
			case "HEIGHT_WO_SHOES":
				record.HeightWoShoes = toPtrFloat(val)
			case "HEIGHT_W_SHOES":
				record.HeightWShoes = toPtrFloat(val)
			case "WEIGHT":
				record.Weight = toPtrFloat(val)
			case "WINGSPAN":
				record.Wingspan = toPtrFloat(val)
			case "STANDING_REACH":
				record.StandingReach = toPtrFloat(val)
			case "BODY_FAT_PCT":
				record.BodyFatPct = toPtrFloat(val)
			case "HAND_LENGTH":
				record.HandLength = toPtrFloat(val)
			case "HAND_WIDTH":
				record.HandWidth = toPtrFloat(val)
			case "STANDING_VERTICAL_LEAP":
				record.StandingVertical = toPtrFloat(val)
			case "MAX_VERTICAL_LEAP":
				record.MaxVertical = toPtrFloat(val)
			case "LANE_AGILITY_TIME":
				record.LaneAgility = toPtrFloat(val)
			case "MODIFIED_LANE_AGILITY_TIME":
				record.ModifiedLaneAgility = toPtrFloat(val)
			case "THREE_QUARTER_SPRINT":
				record.ThreeQuarterSprint = toPtrFloat(val)
			case "BENCH_PRESS":
				record.BenchPress = toPtrInt(val)
			}
		}
		records[i] = record
	}
	return records
}
