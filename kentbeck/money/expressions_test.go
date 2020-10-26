package money

import "testing"

func TestSum(t *testing.T) {
	t.Run("plus returns sum", func(t *testing.T) {
		five := Dollar(5)
		sum := five.Plus(five).(Sum)
		assertMoney(t, five, sum.augend.(Money))
		assertMoney(t, five, sum.addend.(Money))
	})
}
