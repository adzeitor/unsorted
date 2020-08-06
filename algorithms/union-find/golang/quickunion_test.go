package quickfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickUnion_Connected(t *testing.T) {
	t.Skip("Implement QuickFind first")
	// Can be false negative, because it depends on implementation.
	// But it can be very useful for debugging.
	t.Run("step by step (can be false negative)", func(t *testing.T) {
		qu := NewQuickUnion(10)
		assert.Equal(t, QuickUnion{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, qu)

		qu.Union(4, 3)
		assert.Equal(t, QuickUnion{0, 1, 2, 3, 3, 5, 6, 7, 8, 9}, qu)

		qu.Union(3, 8)
		assert.Equal(t, QuickUnion{0, 1, 2, 8, 3, 5, 6, 7, 8, 9}, qu)

		qu.Union(6, 5)
		assert.Equal(t, QuickUnion{0, 1, 2, 8, 3, 5, 5, 7, 8, 9}, qu)

		qu.Union(9, 4)
		assert.Equal(t, QuickUnion{0, 1, 2, 8, 3, 5, 5, 7, 8, 8}, qu)

		qu.Union(2, 1)
		assert.Equal(t, QuickUnion{0, 1, 1, 8, 3, 5, 5, 7, 8, 8}, qu)
		assert.True(t, qu.Connected(8, 9))
		assert.False(t, qu.Connected(5, 4))

		qu.Union(5, 0)
		assert.Equal(t, QuickUnion{0, 1, 1, 8, 3, 0, 5, 7, 8, 8}, qu)

		qu.Union(7, 2)
		assert.Equal(t, QuickUnion{0, 1, 1, 8, 3, 0, 5, 1, 8, 8}, qu)

		qu.Union(6, 1)
		assert.Equal(t, QuickUnion{1, 1, 1, 8, 3, 0, 5, 1, 8, 8}, qu)

		qu.Union(7, 3)
		assert.Equal(t, QuickUnion{1, 8, 1, 8, 3, 0, 5, 1, 8, 8}, qu)
	})

	ConnectivityTests(t, "quick_union_", func(n int) solver {
		return NewQuickUnion(n)
	})
}

func Benchmark_QuickUnion_Union(b *testing.B) {
	benchmarkUnion(b, func(size int) solver {
		return NewQuickUnion(size)
	})
}

func Benchmark_QuickUnion_Connected(b *testing.B) {
	benchmarkConnected(b, func(size int) solver {
		return NewQuickUnion(size)
	})
}
