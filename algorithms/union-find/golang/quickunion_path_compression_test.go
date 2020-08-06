package quickfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickUnionWithPathCompression_New(t *testing.T) {
	t.Run("proper fill initial array", func(t *testing.T) {
		qf := NewQuickUnionWithPathCompression(5)
		assert.Equal(t, []int{0, 1, 2, 3, 4}, qf.ids)
	})
}

func TestQuickUnionWithPathCompression_Connected(t *testing.T) {
	ConnectivityTests(t, "quick_union_path_compression_", func(n int) solver {
		return NewQuickUnionWithPathCompression(n)
	})
}

func TestQuickUnionWithPathCompression_root(t *testing.T) {
	qf := NewQuickUnionWithPathCompression(13)
	// TODO: add ASCII representation
	qf.ids = []int{
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
	assert.Equal(t, qf.ids[9], root)
	assert.Equal(t, qf.ids[6], root)
	assert.Equal(t, qf.ids[3], root)
}
