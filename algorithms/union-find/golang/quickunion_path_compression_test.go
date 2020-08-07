package quickfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickUnionWithPathCompression_Connected(t *testing.T) {
	ConnectivityTests(t, "quick_union_path_compression_", func(n int) solver {
		return NewQuickUnionWithPathCompression(n)
	})
}

func TestQuickUnionWithPathCompression_root(t *testing.T) {
	qf := NewQuickUnionWithPathCompression(13)
	// TODO: add ASCII representation
	qf = []int{
		0:  0, // root
		1:  0,
		2:  0,
		3:  1,
		4:  1,
		5:  1,
		6:  3,
		7:  3,
		8:  6,
		9:  6,
		10: 8,
		11: 9,
		12: 9,
	}
	root := qf.root(9)
	// all visited nodes will be connected to root
	assert.Equal(t, qf[9], root)
	assert.Equal(t, qf[6], root)
	assert.Equal(t, qf[3], root)
}

func Benchmark_QuickUnionWithPathCompression_Union(b *testing.B) {
	benchmarkUnion(b, func(size int) solver {
		return NewQuickUnionWithPathCompression(size)
	})
}

func Benchmark_QuickUnionWithPathCompression_Connected(b *testing.B) {
	benchmarkConnected(b, func(size int) solver {
		return NewQuickUnionWithPathCompression(size)
	})
}
