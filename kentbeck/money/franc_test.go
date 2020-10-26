package money

import (
	"testing"
)

func TestFranc_Times(t *testing.T) {
	five := Franc(5)

	result := five.Times(2)
	assertTrue(t, result.Equal(Franc(10)))

	result = five.Times(3)
	assertTrue(t, result.Equal(Franc(15)))
}

func TestFranc_Equal(t *testing.T) {
	assertTrue(t, Franc(5).Equal(Franc(5)))
	assertFalse(t, Franc(5).Equal(Franc(6)))
}
