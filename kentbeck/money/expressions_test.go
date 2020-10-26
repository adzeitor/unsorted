package money

import (
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("plus returns sum", func(t *testing.T) {
		five := Dollar(5)
		sum := five.Plus(five).(Sum)
		assertMoney(t, five, sum.augend.(Money))
		assertMoney(t, five, sum.addend.(Money))
	})
}

func TestSum_Plus(t *testing.T) {
	t.Run("sum plus money", func(t *testing.T) {
		fiveBucks := Dollar(5)
		tenFrancs := Franc(10)

		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)

		sum := Sum{fiveBucks, tenFrancs}.Plus(fiveBucks)
		result := bank.Reduce(sum, "USD")
		assertMoney(t, Dollar(15), result)
	})

	t.Run("sum times", func(t *testing.T) {
		fiveBucks := Dollar(5)
		tenFrancs := Franc(10)

		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)

		sum := Sum{fiveBucks, tenFrancs}.Times(2)
		result := bank.Reduce(sum, "USD")
		assertMoney(t, Dollar(20), result)
	})
}
