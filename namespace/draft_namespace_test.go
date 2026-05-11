package namespace_test

import (
	"testing"

	"github.com/poteto0/go-nba-sdk/namespace"
	"github.com/stretchr/testify/assert"
)

func Test_CreateDraftNamespace(t *testing.T) {
	// Act
	draftNamespace := namespace.NewDraftNamespace(newProviderForTest())

	// Assert
	assert.NotNil(t, draftNamespace)
}

func Test_Draft_GetCombineStats(t *testing.T) {
	t.Run("can get combine stats", func(t *testing.T) {
		// Arrange
		sl := namespace.NewDraftNamespace(newProviderForTest())

		// Act
		result := sl.GetCombineStats(nil)

		// Assert
		assert.NotNil(t, result.Contents)
	})
}
