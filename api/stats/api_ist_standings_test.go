package stats_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/poteto0/go-nba-sdk/api/stats"
	"github.com/poteto0/go-nba-sdk/constants"
	"github.com/poteto0/go-nba-sdk/fixtures/samples"
	"github.com/poteto0/go-nba-sdk/types"
	"github.com/stretchr/testify/assert"
)

func Test_GetIstStandings(t *testing.T) {
	// Arrange
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	path := constants.StatsBaseUrl + constants.IstStandingsPath + "?LeagueID=00&Season=2023-24"
	httpmock.RegisterResponder(
		"GET",
		path,
		httpmock.NewStringResponder(200, samples.SampleIstStandingsResponse),
	)

	provider := newProviderForTest()

	// Act
	result := stats.GetIstStandings(provider, &types.IstStandingsParams{
		LeagueID: "00",
		Season:   "2023-24",
	})

	// Assert
	assert.Nil(t, result.Error)
	assert.Equal(t, 200, result.StatusCode)
	assert.Equal(t, 1, len(result.Contents.Standings))
	assert.Equal(t, "Lakers", result.Contents.Standings[0].TeamName)
}
