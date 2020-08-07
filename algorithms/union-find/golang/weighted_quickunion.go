package quickfind

// WeightedQuickUnion represent data structure for weighted Quick-union algorithm
// which is  solution to connectivity problem.
type WeightedQuickUnion struct {
	// Value of each index is an index of parent  (ids[i] is parent of i).
	ids     []int
	heights []int
}

func NewWeightedQuickUnion(n int) WeightedQuickUnion {
	ids := make([]int, n)
	heights := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = i
		heights[i] = 1
	}
	return WeightedQuickUnion{
		ids:     ids,
		heights: heights,
	}
}

// Union connects two nodes.
func (qf WeightedQuickUnion) Union(node1, node2 int) {
	root1 := qf.root(node1)
	root2 := qf.root(node2)
	if qf.heights[root1] >= qf.heights[root2] {
		root1, root2 = root2, root1
	}
	qf.ids[root1] = root2
	qf.heights[root2] += qf.heights[root1]
}

// Connected returns true if node1 and node2 are connected.
// In quick-union two nodes is connected if they have the same root.
func (qf WeightedQuickUnion) Connected(node1, node2 int) bool {
	return qf.root(node1) == qf.root(node2)
}

func (qf WeightedQuickUnion) root(node int) int {
	parent := node
	for qf.ids[parent] != parent {
		parent = qf.ids[parent]
	}
	return parent
}
