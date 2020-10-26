package money

import (
	"testing"
)

func TestDollar_Times(t *testing.T) {
	five := Dollar(5)

	result := five.Times(2)
	assert(t, 10, result.Amount)

	result = five.Times(3)
	assert(t, 15, result.Amount)
}

func TestDollar_Equal(t *testing.T) {
	assertTrue(t, Dollar(5).Equal(Dollar(5)))
	assertFalse(t, Dollar(5).Equal(Dollar(6)))
}
