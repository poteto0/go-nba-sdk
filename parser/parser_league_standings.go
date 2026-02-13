package parser

import (
	"github.com/poteto0/go-nba-sdk/types"
)

func ParseLeagueStandingsResponse(rawResponse types.RawResponse) (types.LeagueStandingsResponseContent, error) {
	if rawResponse.Resource != "leaguestandings" {
		return types.LeagueStandingsResponseContent{}, types.NewGnsError(
			"unexpected resource: %s", rawResponse.Resource,
		)
	}

	if len(rawResponse.ResultSets) == 0 {
		return types.LeagueStandingsResponseContent{}, types.NewGnsError(
			"unexpected no record",
		)
	}

	content := types.LeagueStandingsResponseContent{}

	// Extract SeasonYear from parameters if available
	if season, ok := rawResponse.Parameters["Season"].(string); ok {
		content.SeasonYear = season
	}

	for _, resultSet := range rawResponse.ResultSets {
		if resultSet.Name == "Standings" {
			content.Standings = parseLeagueStandingsRecords(resultSet.RowSet, resultSet.Headers)
		}
	}

	return content, nil
}

func parseLeagueStandingsRecords(
	rawSet [][]any, headers []string,
) []types.LeagueStandingsRecord {
	records := make([]types.LeagueStandingsRecord, len(rawSet))
	for i, row := range rawSet {
		var record types.LeagueStandingsRecord
		for j, header := range headers {
			switch header {
			case "LeagueID":
				record.LeagueID = row[j].(string)
			case "SeasonID":
				record.SeasonID = row[j].(string)
			case "TeamID":
				record.TeamID = toInt(row[j])
			case "TeamCity":
				record.TeamCity = row[j].(string)
			case "TeamName":
				record.TeamName = row[j].(string)
			case "Conference":
				record.Conference = row[j].(string)
			case "ConferenceRecord":
				record.ConferenceRecord = row[j].(string)
			case "PlayoffRank":
				record.PlayoffRank = toInt(row[j])
			case "ClinchIndicator":
				record.ClinchIndicator = row[j].(string)
			case "Division":
				record.Division = row[j].(string)
			case "DivisionRecord":
				record.DivisionRecord = row[j].(string)
			case "DivisionRank":
				record.DivisionRank = toInt(row[j])
			case "WINS":
				record.Wins = toInt(row[j])
			case "LOSSES":
				record.Losses = toInt(row[j])
			case "WinPCT":
				record.WinPct = toFloat(row[j])
			case "LeagueRank":
				record.LeagueRank = toInt(row[j])
			case "Record":
				record.Record = row[j].(string)
			case "HOME":
				record.Home = row[j].(string)
			case "ROAD":
				record.Road = row[j].(string)
			case "L10":
				record.L10 = row[j].(string)
			case "Last10Home":
				record.Last10Home = row[j].(string)
			case "Last10Road":
				record.Last10Road = row[j].(string)
			case "OT":
				record.OT = row[j].(string)
			case "ThreePTSOrLess":
				record.ThreePTSOrLess = row[j].(string)
			case "TenPTSOrMore":
				record.TenPTSOrMore = row[j].(string)
			case "LongHomeStreak":
				record.LongHomeStreak = toInt(row[j])
			case "strLongHomeStreak":
				record.StrLongHomeStreak = row[j].(string)
			case "LongRoadStreak":
				record.LongRoadStreak = toInt(row[j])
			case "strLongRoadStreak":
				record.StrLongRoadStreak = row[j].(string)
			case "LongWinStreak":
				record.LongWinStreak = toInt(row[j])
			case "LongLossStreak":
				record.LongLossStreak = toInt(row[j])
			case "CurrentHomeStreak":
				record.CurrentHomeStreak = toInt(row[j])
			case "strCurrentHomeStreak":
				record.StrCurrentHomeStreak = row[j].(string)
			case "CurrentRoadStreak":
				record.CurrentRoadStreak = toInt(row[j])
			case "strCurrentRoadStreak":
				record.StrCurrentRoadStreak = row[j].(string)
			case "CurrentStreak":
				record.CurrentStreak = toInt(row[j])
			case "strCurrentStreak":
				record.StrCurrentStreak = row[j].(string)
			case "ConferenceGamesBack":
				record.ConferenceGamesBack = toFloat(row[j])
			case "DivisionGamesBack":
				record.DivisionGamesBack = toFloat(row[j])
			case "ClinchedConferenceTitle":
				record.ClinchedConferenceTitle = toPtrString(row[j])
			case "ClinchedDivisionTitle":
				record.ClinchedDivisionTitle = toPtrString(row[j])
			case "ClinchedPlayoffBirth":
				record.ClinchedPlayoffBirth = toPtrString(row[j])
			case "EliminatedConference":
				record.EliminatedConference = toPtrString(row[j])
			case "EliminatedDivision":
				record.EliminatedDivision = toPtrString(row[j])
			case "AheadAtHalf":
				record.AheadAtHalf = row[j].(string)
			case "BehindAtHalf":
				record.BehindAtHalf = row[j].(string)
			case "TiedAtHalf":
				record.TiedAtHalf = row[j].(string)
			case "AheadAtThird":
				record.AheadAtThird = row[j].(string)
			case "BehindAtThird":
				record.BehindAtThird = row[j].(string)
			case "TiedAtThird":
				record.TiedAtThird = row[j].(string)
			case "Score100PTS":
				record.Score100PTS = row[j].(string)
			case "OppScore100PTS":
				record.OppScore100PTS = row[j].(string)
			case "OppOver500":
				record.OppOver500 = row[j].(string)
			case "LeadInFGPCT":
				record.LeadInFGPCT = row[j].(string)
			case "LeadInReb":
				record.LeadInReb = row[j].(string)
			case "FewerTurnovers":
				record.FewerTurnovers = row[j].(string)
			case "PointsPG":
				record.PointsPG = toFloat(row[j])
			case "OppPointsPG":
				record.OppPointsPG = toFloat(row[j])
			case "DiffPointsPG":
				record.DiffPointsPG = toFloat(row[j])
			case "vsEast":
				record.VsEast = row[j].(string)
			case "vsAtlantic":
				record.VsAtlantic = row[j].(string)
			case "vsCentral":
				record.VsCentral = row[j].(string)
			case "vsSoutheast":
				record.VsSoutheast = row[j].(string)
			case "vsWest":
				record.VsWest = row[j].(string)
			case "vsNorthwest":
				record.VsNorthwest = row[j].(string)
			case "vsPacific":
				record.VsPacific = row[j].(string)
			case "vsSouthwest":
				record.VsSouthwest = row[j].(string)
			case "Jan":
				record.Jan = toPtrString(row[j])
			case "Feb":
				record.Feb = toPtrString(row[j])
			case "Mar":
				record.Mar = toPtrString(row[j])
			case "Apr":
				record.Apr = toPtrString(row[j])
			case "May":
				record.May = toPtrString(row[j])
			case "Jun":
				record.Jun = toPtrString(row[j])
			case "Jul":
				record.Jul = toPtrString(row[j])
			case "Aug":
				record.Aug = toPtrString(row[j])
			case "Sep":
				record.Sep = toPtrString(row[j])
			case "Oct":
				record.Oct = toPtrString(row[j])
			case "Nov":
				record.Nov = toPtrString(row[j])
			case "Dec":
				record.Dec = toPtrString(row[j])
			case "PreAS":
				record.PreAS = toPtrString(row[j])
			case "PostAS":
				record.PostAS = toPtrString(row[j])
			}
		}
		records[i] = record
	}

	return records
}
