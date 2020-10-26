package money

import (
	"testing"
)

func TestDollar_Times(t *testing.T) {
	five := Dollar(5)

	result := five.Times(2)
	assertEqual(t, Dollar(10), result)

	result = five.Times(3)
	assertEqual(t, Dollar(15), result)
}

func TestDollar_Equal(t *testing.T) {
	assertTrue(t, Dollar(5).Equal(Dollar(5)))
	assertFalse(t, Dollar(5).Equal(Dollar(6)))
}
