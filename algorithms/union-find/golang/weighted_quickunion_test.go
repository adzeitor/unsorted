package quickfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeightedQuickUnion_Connected(t *testing.T) {
	t.Skip("Implement QuickUnion first")
	// Can be false negative, because it depends on implementation.
	// But it can be very useful for debugging.
	t.Run("step by step (can be false negative)", func(t *testing.T) {
		qu := NewWeightedQuickUnion(10)
		assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, qu.ids)

		qu.Union(4, 3)
		assert.Equal(t, []int{0, 1, 2, 4, 4, 5, 6, 7, 8, 9}, qu.ids)

		qu.Union(3, 8)
		assert.Equal(t, []int{0, 1, 2, 4, 4, 5, 6, 7, 4, 9}, qu.ids)

		qu.Union(6, 5)
		assert.Equal(t, []int{0, 1, 2, 4, 4, 6, 6, 7, 4, 9}, qu.ids)

		qu.Union(9, 4)
		assert.Equal(t, []int{0, 1, 2, 4, 4, 6, 6, 7, 4, 4}, qu.ids)

		qu.Union(2, 1)
		assert.Equal(t, []int{0, 2, 2, 4, 4, 6, 6, 7, 4, 4}, qu.ids)
		assert.True(t, qu.Connected(8, 9))
		assert.False(t, qu.Connected(5, 4))

		qu.Union(5, 0)
		assert.Equal(t, []int{6, 2, 2, 4, 4, 6, 6, 7, 4, 4}, qu.ids)

		qu.Union(7, 2)
		assert.Equal(t, []int{6, 2, 2, 4, 4, 6, 6, 2, 4, 4}, qu.ids)

		qu.Union(6, 1)
		assert.Equal(t, []int{6, 2, 6, 4, 4, 6, 6, 2, 4, 4}, qu.ids)

		qu.Union(7, 3)
		assert.Equal(t, []int{6, 2, 6, 4, 6, 6, 6, 2, 4, 4}, qu.ids)
	})

	ConnectivityTests(t, "weighted_quick_union_", func(n int) solver {
		return NewWeightedQuickUnion(n)
	})
}

func Benchmark_WeightedQuickUnion_Union(b *testing.B) {
	benchmarkUnion(b, func(size int) solver {
		return NewWeightedQuickUnion(size)
	})
}

func Benchmark_WeightedQuickUnion_Connected(b *testing.B) {
	benchmarkConnected(b, func(size int) solver {
		return NewWeightedQuickUnion(size)
	})
}
