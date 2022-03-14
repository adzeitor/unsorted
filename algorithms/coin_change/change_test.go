package coin_change

import (
	"reflect"
	"testing"
)

func TestMakeChange(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		assert(t, 5, MakeChange(10, []int{2, 5, 3, 6}))
		assert(t, 13, MakeChange(27, []int{25, 10, 5, 1}))
		assert(t, 134, MakeChange(79, []int{50, 25, 10, 5, 1}))
		assert(t, 2, MakeChange(27, []int{25, 1}))
		assert(t, 75, MakeChange(27, []int{1, 2, 3}))
		assert(t, 6, MakeChange(7, []int{1, 2, 5}))
		assert(t, 3003001, MakeChange(6000, []int{1, 2, 3}))
	})

	t.Run("no answer", func(t *testing.T) {
		assert(t, 0, MakeChange(10, []int{15}))
		assert(t, 0, MakeChange(5, []int{}))
	})
}

func assert(t *testing.T, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("not equal want=%+v got=%+v", want, got)
	}
}
