package quickfind

// QuickUnion represent data structure for Quick-union algorithm which is
// solution to connectivity problem.
// Value of each index is an index of parent  (id[i] is parent of i).
type QuickUnion []int

func NewQuickUnion(n int) QuickUnion {
	panic("implement me")
}

// Union connects two nodes.
// To merge components containing node1 and node2, set the id of node1's root
// to the id of node2's root.
func (qf QuickUnion) Union(node1, node2 int) {
	panic("implement me")
}

// Connected returns true if node1 and node2 are connected.
// In quick-union two nodes is connected if they have the same root.
func (qf QuickUnion) Connected(node1, node2 int) bool {
	panic("implement me")
}
