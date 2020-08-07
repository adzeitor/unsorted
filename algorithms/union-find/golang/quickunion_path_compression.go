package quickfind

// QuickUnionWithPathCompression represent data structure for weighted Quick-union algorithm
// which is  solution to connectivity problem.
type QuickUnionWithPathCompression []int

func NewQuickUnionWithPathCompression(n int) QuickUnionWithPathCompression {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
	}
	return ids
}

// Union connects two nodes.
func (qf QuickUnionWithPathCompression) Union(node1, node2 int) {
	qf[qf.root(node1)] = qf.root(node2)
}

// Connected returns true if node1 and node2 are connected.
// In quick-union two nodes is connected if they have the same root.
func (qf QuickUnionWithPathCompression) Connected(node1, node2 int) bool {
	return qf.root(node1) == qf.root(node2)
}

func (qf QuickUnionWithPathCompression) root(node int) int {
	var visited []int
	root := node
	for qf[root] != root {
		visited = append(visited, root)
		root = qf[root]
	}

	for _, id := range visited {
		qf[id] = root
	}
	return root
}
