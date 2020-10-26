package money

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("not equal %s %s", want, got)
	}
}
