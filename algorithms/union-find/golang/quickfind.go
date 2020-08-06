package quickfind

// QuickFind represents data structure for Quick-find algorithm which is
// solution to connectivity problem.
// Nodes (indices) are connected if they have the same values.
type QuickFind []int

func NewQuickFind(n int) QuickFind {
	panic("implement me")
}

func (qf QuickFind) Union(node1, node2 int) {
	panic("implement me")
}

func (qf QuickFind) Connected(node1, node2 int) bool {
	panic("implement me")
}
