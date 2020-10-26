package money

import (
	"reflect"
	"testing"
)

func assert(t *testing.T, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("not equal want=%+v got=%+v", want, got)
	}
}

func assertTrue(t *testing.T, got interface{}) {
	t.Helper()
	assert(t, true, got)
}

func assertFalse(t *testing.T, got interface{}) {
	t.Helper()
	assert(t, false, got)
}

type Equaler interface {
	Equal(other interface{}) bool
}

func assertEqual(t *testing.T, want, got Equaler) {
	assertTrue(t, want.Equal(got))
}
