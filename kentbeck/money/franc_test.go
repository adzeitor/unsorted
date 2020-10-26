package money

import (
	"testing"
)

func TestFranc_Times(t *testing.T) {
	five := Franc(5)

	assertEqual(t, Franc(10), five.Times(2))
	assertEqual(t, Franc(15), five.Times(3))
}

func TestFranc_Equal(t *testing.T) {
	assertTrue(t, Franc(5).Equal(Franc(5)))
	assertFalse(t, Franc(5).Equal(Franc(6)))
}
