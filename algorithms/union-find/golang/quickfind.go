package quickfind

// QuickFind represents data structure for Quick-find algorithm which is
// solution to connectivity problem.
// Nodes (indices) are connected if they have the same values.
type QuickFind []int

func NewQuickFind(n int) QuickFind {
	qf := make(QuickFind, n)
	for i := 0; i < n; i++ {
		qf[i] = i
	}
	return qf
}

func (qf QuickFind) Union(node1, node2 int) {
	old := qf[node1]
	for i := range qf {
		if qf[i] == old {
			qf[i] = qf[node2]
		}
	}
}

func (qf QuickFind) Connected(node1, node2 int) bool {
	return qf[node1] == qf[node2]
}
