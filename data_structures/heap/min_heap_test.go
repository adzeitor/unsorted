package heap

import (
	"reflect"
	"testing"
)

func TestMinHeap(t *testing.T) {
	t.Run("Simple case", func(t *testing.T) {
		// arrange
		heap := NewMinHeap(func(one int, other int) bool { return one < other })

		// act
		heap.Add(5)
		heap.Add(2)
		heap.Add(1)
		heap.Add(3)
		heap.Add(4)
		heap.Add(6)

		t.Log(heap.elements)
		// assert
		assert(t, 1, heap.Remove())
		assert(t, 2, heap.Remove())
		assert(t, 3, heap.Remove())
		assert(t, 4, heap.Remove())
		assert(t, 5, heap.Remove())
		assert(t, 6, heap.Remove())
	})

	t.Run("Add after removing", func(t *testing.T) {
		// arrange
		heap := NewMinHeap(func(one string, other string) bool { return one < other })
		heap.Add("a")
		heap.Add("b")
		heap.Add("c")
		heap.Add("d")
		heap.Add("e")
		heap.Add("f")

		// act
		heap.Remove()
		heap.Add("a")

		// assert
		assert(t, "a", heap.Remove())
		assert(t, "b", heap.Remove())
		assert(t, "c", heap.Remove())
	})
}

func assert(t *testing.T, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("not equal want=%+v got=%+v", want, got)
	}
}
