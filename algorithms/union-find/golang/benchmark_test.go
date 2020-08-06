package quickfind

import (
	"fmt"
	"math/rand"
	"testing"
)

func fillBenchmarkData(solver solver, size int, connections int) {
	for i := 0; i < connections; i++ {
		node1 := rand.Intn(size)
		node2 := rand.Intn(size)
		solver.Union(node1, node2)
	}
}

func benchmarkConnected(b *testing.B, factory func(size int) solver) {
	sizes := []int{10, 100, 1000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			solver := factory(size)
			connections := size / 2
			fillBenchmarkData(solver, size, connections)

			b.StartTimer()
			for n := 0; n < b.N; n++ {
				for node := 1; node < size; node++ {
					solver.Connected(node, node-1)
				}
			}
			b.StopTimer()
		})
	}
}

func Benchmark_QuickFind_Connected(b *testing.B) {
	benchmarkConnected(b, func(size int) solver {
		return NewQuickFind(size)
	})
}

func Benchmark_QuickUnion_Connected(b *testing.B) {
	benchmarkConnected(b, func(size int) solver {
		return NewQuickUnion(size)
	})
}

func Benchmark_WeightedQuickUnion_Connected(b *testing.B) {
	benchmarkConnected(b, func(size int) solver {
		return NewWeightedQuickUnion(size)
	})
}

func Benchmark_QuickUnionWithPathCompression_Connected(b *testing.B) {
	benchmarkConnected(b, func(size int) solver {
		return NewQuickUnionWithPathCompression(size)
	})
}

func benchmarkUnion(b *testing.B, factory func(size int) solver) {
	sizes := []int{10, 100, 1000, 10000, 100000, 1_000_000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
			solver := factory(size)
			connections := size / 2

			b.StartTimer()
			for n := 0; n < b.N; n++ {
				for i := 1; i < connections; i++ {
					node1 := i
					node2 := i - 1
					solver.Union(node1, node2)
				}
			}
			b.StopTimer()
		})
	}
}

func Benchmark_QuickFind_Union(b *testing.B) {
	benchmarkUnion(b, func(size int) solver {
		return NewQuickFind(size)
	})
}

func Benchmark_QuickUnion_Union(b *testing.B) {
	benchmarkUnion(b, func(size int) solver {
		return NewQuickUnion(size)
	})
}

func Benchmark_WeightedQuickUnion_Union(b *testing.B) {
	benchmarkUnion(b, func(size int) solver {
		return NewWeightedQuickUnion(size)
	})
}

func Benchmark_QuickUnionWithPathCompression_Union(b *testing.B) {
	benchmarkUnion(b, func(size int) solver {
		return NewQuickUnionWithPathCompression(size)
	})
}
