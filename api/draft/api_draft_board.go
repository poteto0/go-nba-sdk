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

func GetDraftBoard(provider api.IProvider, params *types.DraftBoardParams) types.Response[types.DraftBoardResponse] {
	if params == nil {
		params = &types.DraftBoardParams{}
	}

	if params.LeagueID == "" {
		params.LeagueID = "00"
	}

	path := constants.StatsBaseUrl + constants.DraftBoardPath
	v, err := query.Values(params)
	if err != nil {
		return types.Response[types.DraftBoardResponse]{Error: err}
	}

	path = path + "?" + v.Encode()

	resp, err := provider.Get(path, &constants.DefaultStatsHeaders)
	if err != nil {
		return types.Response[types.DraftBoardResponse]{Error: err}
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return types.Response[types.DraftBoardResponse]{StatusCode: resp.StatusCode, Error: err}
	}
	if len(bodyBytes) == 0 {
		return types.Response[types.DraftBoardResponse]{StatusCode: resp.StatusCode, Error: types.NewGnsError("empty body. Status: %d", resp.StatusCode)}
	}

	var contents types.DraftBoardResponse
	err = json.Unmarshal(bodyBytes, &contents)
	if err != nil {
		return types.Response[types.DraftBoardResponse]{
			StatusCode: resp.StatusCode,
			Error:      fmt.Errorf("json unmarshal err: %w, body: %s", err, string(bodyBytes)),
		}
	}

	return types.Response[types.DraftBoardResponse]{
		Contents:   contents,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
