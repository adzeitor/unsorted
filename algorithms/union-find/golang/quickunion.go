package quickfind

// QuickUnion represent data structure for Quick-union algorithm which is
// solution to connectivity problem.
// Value of each index is an index of parent  (id[i] is parent of i).
type QuickUnion []int

func NewQuickUnion(n int) QuickUnion {
	qf := make(QuickUnion, n)
	for i := 0; i < n; i++ {
		qf[i] = i
	}
	return qf
}

// Union connects two nodes.
// To merge components containing node1 and node2, set the id of node1's root
// to the id of node2's root.
func (qf QuickUnion) Union(node1, node2 int) {
	qf[qf.root(node1)] = qf.root(node2)
}

// Connected returns true if node1 and node2 are connected.
// In quick-union two nodes is connected if they have the same root.
func (qf QuickUnion) Connected(node1, node2 int) bool {
	return qf.root(node1) == qf.root(node2)
}

func (qf QuickUnion) root(node int) int {
	parent := node
	for qf[parent] != parent {
		parent = qf[parent]
	}
	return parent
}
