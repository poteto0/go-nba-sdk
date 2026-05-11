package draft_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/api"
	"github.com/poteto0/go-nba-sdk/api/draft"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func newProviderForTest() api.IProvider {
	return api.NewProvider(nil)
}

func Test_GetDraftBoard(t *testing.T) {
	t.Run("can get draft board", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.DraftBoardPath,
			httpmock.NewStringResponder(200, samples.SampleDraftBoardResponse),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := draft.GetDraftBoard(provider, &types.DraftBoardParams{
			LeagueID:   "00",
			SeasonYear: "2025",
		})

		// Assert
		assert.NoError(t, result.Error)
		assert.NotNil(t, result.Contents)
		assert.Equal(t, 200, result.StatusCode)

		// assert contents
		assert.Equal(t, "draftboard", *result.Contents.Resource)
		assert.NotEmpty(t, result.Contents.ResultSets)
	})

	t.Run("network error is w/o status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.DraftBoardPath,
			httpmock.NewErrorResponder(assert.AnError),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := draft.GetDraftBoard(provider, &types.DraftBoardParams{
			LeagueID:   "00",
			SeasonYear: "2025",
		})

		// Assert
		assert.Error(t, result.Error)
		assert.Equal(t, 0, result.StatusCode)
	})

	t.Run("response parse error w/ http status code", func(t *testing.T) {
		httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			constants.StatsBaseUrl+constants.DraftBoardPath,
			httpmock.NewStringResponder(200, "invalid"),
		)

		// Arrange
		provider := newProviderForTest()

		// Act
		result := draft.GetDraftBoard(provider, &types.DraftBoardParams{
			LeagueID:   "00",
			SeasonYear: "2025",
		})

		// Assert
		assert.Equal(t, 200, result.StatusCode)
		assert.Error(t, result.Error)
	})
}
