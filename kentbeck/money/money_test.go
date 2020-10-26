package money

import (
	"testing"
)

func TestDollar_Times(t *testing.T) {
	five := Dollar(5)

	result := five.Times(2)
	assertTrue(t, result.Equal(Dollar(10)))

	result = five.Times(3)
	assertTrue(t, result.Equal(Dollar(15)))
}

func TestDollar_Equal(t *testing.T) {
	assertTrue(t, Dollar(5).Equal(Dollar(5)))
	assertFalse(t, Dollar(5).Equal(Dollar(6)))
}
