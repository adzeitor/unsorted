package quickfind

// QuickUnionWithPathCompression represent data structure for weighted Quick-union algorithm
// which is  solution to connectivity problem.
type QuickUnionWithPathCompression []int

func NewQuickUnionWithPathCompression(n int) QuickUnionWithPathCompression {
	panic("implement me")
}

// Union connects two nodes.
func (qf QuickUnionWithPathCompression) Union(node1, node2 int) {
	panic("implement me")
}

// Connected returns true if node1 and node2 are connected.
// In quick-union two nodes is connected if they have the same root.
func (qf QuickUnionWithPathCompression) Connected(node1, node2 int) bool {
	panic("implement me")
}

func (qf QuickUnionWithPathCompression) root(node int) int {
	panic("implement me")
}
