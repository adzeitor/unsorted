package money

import (
	"testing"
)

func TestDollar_Times(t *testing.T) {
	five := Dollar(5)
	five.Times(2)
	assert(t, 10, five.Amount)
}
