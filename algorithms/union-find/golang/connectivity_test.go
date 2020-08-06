package quickfind

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type solver interface {
	Union(node1, node2 int)
	Connected(node1, node2 int) bool
}

func ConnectivityTests(t *testing.T, name string, factory func(n int) solver) {
	t.Run(name+"each node is connected with itself", func(t *testing.T) {
		qf := factory(5)
		assert.True(t, qf.Connected(0, 0))
		assert.True(t, qf.Connected(1, 1))
		assert.True(t, qf.Connected(2, 2))
		assert.True(t, qf.Connected(3, 3))
		assert.True(t, qf.Connected(4, 4))
	})

	t.Run(name+"nodes in empty structure are not connected", func(t *testing.T) {
		qf := factory(5)
		assert.False(t, qf.Connected(0, 1))
		assert.False(t, qf.Connected(1, 2))
		assert.False(t, qf.Connected(2, 3))
		assert.False(t, qf.Connected(3, 4))
	})

	t.Run(name+"union nodes are connected", func(t *testing.T) {
		qf := factory(5)
		qf.Union(1, 2)
		qf.Union(3, 4)

		assert.True(t, qf.Connected(1, 2), qf)
		assert.True(t, qf.Connected(3, 4), qf)
		// other nodes still not connected
		assert.False(t, qf.Connected(0, 1))
		assert.False(t, qf.Connected(1, 3))
		assert.False(t, qf.Connected(2, 3))
		assert.False(t, qf.Connected(4, 2))
	})

	t.Run(name+"example from MOOC", func(t *testing.T) {
		qf := factory(10)
		qf.Union(3, 4)
		qf.Union(3, 8)
		qf.Union(5, 6)
		qf.Union(4, 9)
		qf.Union(1, 2)
		qf.Union(0, 5)
		qf.Union(2, 7)
		qf.Union(1, 6)

		// connectivity will be:
		//   0 - 5 - 6 - 1 - 2
		//   8 - 3 - 4 - 9
		assert.True(t, qf.Connected(0, 5))
		assert.True(t, qf.Connected(5, 6))
		assert.True(t, qf.Connected(6, 1))
		assert.True(t, qf.Connected(1, 2))

		assert.True(t, qf.Connected(8, 3))
		assert.True(t, qf.Connected(3, 4))
		assert.True(t, qf.Connected(4, 9))

		// two clusters are not connected
		assert.False(t, qf.Connected(8, 2))
		assert.False(t, qf.Connected(5, 4))
	})

	t.Run(name+"example 2 from MOOC (all connected)", func(t *testing.T) {
		qf := factory(10)
		qf.Union(3, 4)
		qf.Union(3, 8)
		qf.Union(5, 6)
		qf.Union(4, 9)
		qf.Union(1, 2)
		qf.Union(0, 5)
		qf.Union(2, 7)
		qf.Union(1, 6)
		qf.Union(7, 3)

		assert.True(t, qf.Connected(0, 5))
		assert.True(t, qf.Connected(5, 6))
		assert.True(t, qf.Connected(6, 1))
		assert.True(t, qf.Connected(1, 2))
		assert.True(t, qf.Connected(8, 3))
		assert.True(t, qf.Connected(3, 4))
		assert.True(t, qf.Connected(4, 9))
		assert.True(t, qf.Connected(8, 2))
		assert.True(t, qf.Connected(5, 4))
		t.Log(qf)
	})

	t.Run(name+"all nodes are connected", func(t *testing.T) {
		qf := factory(4)
		qf.Union(3, 2)
		qf.Union(2, 1)
		qf.Union(1, 0)

		assert.True(t, qf.Connected(0, 0), qf)
		assert.True(t, qf.Connected(1, 1), qf)
		assert.True(t, qf.Connected(2, 2), qf)
		assert.True(t, qf.Connected(3, 3), qf)

		assert.True(t, qf.Connected(0, 1), qf)
		assert.True(t, qf.Connected(0, 2), qf)
		assert.True(t, qf.Connected(0, 3), qf)
		assert.True(t, qf.Connected(1, 2), qf)
		assert.True(t, qf.Connected(1, 3), qf)
		assert.True(t, qf.Connected(2, 3), qf)
	})
}
