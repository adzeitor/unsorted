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
