package live_test

import "github.com/poteto0/go-nba-sdk/api"

func newProviderForTest() api.IProvider {
	return api.NewProvider(nil)
}
