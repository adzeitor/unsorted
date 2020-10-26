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

func assertMoney(t *testing.T, want, got Money) {
	t.Helper()
	if !reflect.DeepEqual(want.Equal(got), true) {
		t.Errorf("not equal want=%+v got=%+v", want, got)
	}
}
