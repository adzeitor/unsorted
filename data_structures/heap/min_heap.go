package heap

type MinHeap[T any] struct {
	//
	//          1
	//     2         3
	//  4    7     6   5
	//
	// will be represented as array:
	// 1 2 3 4 7 6 5
	//
	elements []T
	less     func(T, T) bool
}

func NewMinHeap[T any](lessCmp func(T, T) bool) *MinHeap[T] {
	return &MinHeap[T]{
		elements: nil,
		less:     lessCmp,
	}
}

// Add value to heap.
func (heap *MinHeap[T]) Add(value T) {
	heap.elements = append(heap.elements, value)
	heap.heapifyUp(heap.lastElement())
}

// Remove returns min element from heap.
func (heap *MinHeap[T]) Remove() T {
	if len(heap.elements) == 0 {
		panic("heap is empty")
	}
	top := heap.elements[0]
	heap.elements[0] = heap.elements[heap.lastElement()]
	heap.elements = heap.elements[:heap.lastElement()]
	heap.heapifyDown(0)
	return top
}

// lastElement returns index of last element.
func (heap *MinHeap[T]) lastElement() int {
	return len(heap.elements) - 1
}

func (heap *MinHeap[T]) heapifyUp(i int) {
	// FIXME: special case? needed?
	if i == 0 {
		return
	}

	parentIndex := heap.parent(i)
	if heap.less(heap.elements[i], heap.elements[parentIndex]) {
		heap.swap(parentIndex, i)
		heap.heapifyUp(parentIndex)
	}
}

func (heap *MinHeap[T]) heapifyDown(i int) {
	left := heap.leftChild(i)
	if left != -1 && heap.less(heap.elements[left], heap.elements[i]) {
		heap.swap(left, i)
		heap.heapifyDown(i)
		return
	}

	right := heap.rightChild(i)
	if right != -1 && heap.less(heap.elements[right], heap.elements[i]) {
		heap.swap(right, i)
		heap.heapifyDown(i)
		return
	}
}

// parent returns index of parent element of specified child.
func (heap *MinHeap[T]) parent(child int) int {
	return (child - 1) / 2
}

func (heap *MinHeap[T]) leftChild(parentIndex int) int {
	index := 2*parentIndex + 1
	if index >= len(heap.elements) {
		return -1
	}
	return index
}

func (heap *MinHeap[T]) rightChild(parentIndex int) int {
	index := 2*parentIndex + 2
	if index >= len(heap.elements) {
		return -1
	}
	return index
}

// swap elements of specified indices.
func (heap *MinHeap[T]) swap(i, j int) {
	heap.elements[i], heap.elements[j] = heap.elements[j], heap.elements[i]
}
