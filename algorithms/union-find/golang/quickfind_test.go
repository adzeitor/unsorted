package quickfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickFind_New(t *testing.T) {
	t.Run("proper fill initial array", func(t *testing.T) {
		qf := NewQuickFind(5)
		assert.Equal(t, QuickFind{0, 1, 2, 3, 4}, qf)
	})
}

func TestQuickFind_Connected(t *testing.T) {
	// Can be false negative, because it depends on implementation.
	// But it can be very useful for debugging.
	t.Run("step by step (can be false negative)", func(t *testing.T) {
		qu := NewQuickFind(10)
		assert.Equal(t, QuickFind{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, qu)

		qu.Union(4, 3)
		assert.Equal(t, QuickFind{0, 1, 2, 3, 3, 5, 6, 7, 8, 9}, qu)

		qu.Union(3, 8)
		assert.Equal(t, QuickFind{0, 1, 2, 8, 8, 5, 6, 7, 8, 9}, qu)

		qu.Union(6, 5)
		assert.Equal(t, QuickFind{0, 1, 2, 8, 8, 5, 5, 7, 8, 9}, qu)

		qu.Union(9, 4)
		assert.Equal(t, QuickFind{0, 1, 2, 8, 8, 5, 5, 7, 8, 8}, qu)

		qu.Union(2, 1)
		assert.Equal(t, QuickFind{0, 1, 1, 8, 8, 5, 5, 7, 8, 8}, qu)
		assert.True(t, qu.Connected(8, 9))
		assert.False(t, qu.Connected(5, 0))

		qu.Union(5, 0)
		assert.Equal(t, QuickFind{0, 1, 1, 8, 8, 0, 0, 7, 8, 8}, qu)

		qu.Union(7, 2)
		assert.Equal(t, QuickFind{0, 1, 1, 8, 8, 0, 0, 1, 8, 8}, qu)

		qu.Union(6, 1)
		assert.Equal(t, QuickFind{1, 1, 1, 8, 8, 1, 1, 1, 8, 8}, qu)
	})

	ConnectivityTests(t, "quick_find_", func(n int) solver {
		return NewQuickFind(n)
	})
}
