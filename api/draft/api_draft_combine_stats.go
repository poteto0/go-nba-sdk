package draft

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/google/go-querystring/query"
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/types"
)

func GetCombineStats(provider api.IProvider, params *types.DraftCombineStatsParams) types.Response[types.DraftCombineStatsResponse] {
	if params == nil {
		params = &types.DraftCombineStatsParams{}
	}

	if params.LeagueID == "" {
		params.LeagueID = "00"
	}

	path := constants.StatsBaseUrl + constants.DraftCombineStatsPath
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.DraftCombineStatsResponse]{Error: err}
	}

	path = path + "?" + v.Encode()

	resp, err := provider.Get(path, &constants.DefaultStatsHeaders)
	if err != nil {
		return types.Response[types.DraftCombineStatsResponse]{Error: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Response[types.DraftCombineStatsResponse]{StatusCode: resp.StatusCode, Error: err}
	}
	if len(bodyBytes) == 0 {
		return types.Response[types.DraftCombineStatsResponse]{StatusCode: resp.StatusCode, Error: types.NewGnsError("empty body. Status: %d", resp.StatusCode)}
	}

	var contents types.DraftCombineStatsResponse
	err = json.Unmarshal(bodyBytes, &contents)
	if err != nil {
		return types.Response[types.DraftCombineStatsResponse]{
			StatusCode: resp.StatusCode,
			Error:      fmt.Errorf("json unmarshal err: %w, body: %s", err, string(bodyBytes)),
		}
	}

	return types.Response[types.DraftCombineStatsResponse]{
		Contents:   contents,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
