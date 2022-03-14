package binary_search_tree

import (
	"reflect"
	"testing"
)

func TestTree(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tree := New()

		tree.Insert(5)
		tree.Insert(4)
		tree.Insert(6)

		// found
		assert(t, 4, *tree.Search(4))
		assert(t, 5, *tree.Search(5))
		assert(t, 6, *tree.Search(6))

		// not found
		assert(t, true, nil == tree.Search(0))
		assert(t, true, nil == tree.Search(3))
		assert(t, true, nil == tree.Search(7))
	})
}

func assert(t *testing.T, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("not equal want=%+v got=%+v", want, got)
	}
}
