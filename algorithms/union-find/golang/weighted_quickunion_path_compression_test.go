package quickfind

import (
	"testing"
)

func TestWeightedQuickUnionWithPathCompression_Connected(t *testing.T) {
	t.Skip("Implement WeightedQuickUnion first")
	ConnectivityTests(t, "weighted_quick_union_", func(n int) solver {
		return NewWeightedQuickUnionWithPathCompression(n)
	})
}

func Benchmark_WeightedQuickUnionWithPathCompression_Union(b *testing.B) {
	benchmarkUnion(b, func(size int) solver {
		return NewWeightedQuickUnionWithPathCompression(size)
	})
}

func Benchmark_WeightedQuickUnionWithPathCompression_Connected(b *testing.B) {
	benchmarkConnected(b, func(size int) solver {
		return NewWeightedQuickUnionWithPathCompression(size)
	})
}
