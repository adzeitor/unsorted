package quickfind

// WeightedQuickUnion represent data structure for weighted Quick-union algorithm
// which is  solution to connectivity problem.
type WeightedQuickUnion struct {
	// Value of each index is an index of parent  (ids[i] is parent of i).
	ids     []int
	heights []int
}

func NewWeightedQuickUnion(n int) WeightedQuickUnion {
	panic("implement me")
}

// Union connects two nodes.
func (qf WeightedQuickUnion) Union(node1, node2 int) {
	panic("implement me")
}

// Connected returns true if node1 and node2 are connected.
// In quick-union two nodes is connected if they have the same root.
func (qf WeightedQuickUnion) Connected(node1, node2 int) bool {
	panic("implement me")
}
