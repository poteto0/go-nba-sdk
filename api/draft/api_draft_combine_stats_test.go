package draft_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/api/draft"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_GetCombineStats(t *testing.T) {
	t.Run("can get combine stats", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.DraftCombineStatsPath,
			httpmock.NewStringResponder(200, samples.SampleDraftCombineStatsResponse),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := draft.GetCombineStats(provider, &types.DraftCombineStatsParams{
			LeagueID:   "00",
			SeasonYear: "2025-26",
		})

		// Assert
		assert.NoError(t, result.Error)
		assert.NotNil(t, result.Contents)
		assert.Equal(t, 200, result.StatusCode)
		
		// assert contents
		assert.Equal(t, "draftcombinestats", *result.Contents.Resource)
		assert.NotEmpty(t, result.Contents.ResultSets)
	})
}
