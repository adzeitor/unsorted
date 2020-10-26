package money

import (
	"testing"
)

func TestBank_Reduce(t *testing.T) {
	t.Run("reduce sum", func(t *testing.T) {
		sum := Dollar(3).Plus(Dollar(4))
		bank := NewBank()

		reduced := bank.Reduce(sum, "USD")
		assertMoney(t, Dollar(7), reduced)
	})

	t.Run("reduce money", func(t *testing.T) {
		bank := NewBank()
		reduced := bank.Reduce(Dollar(1), "USD")
		assertMoney(t, Dollar(1), reduced)
	})

	t.Run("reduce money different currency", func(t *testing.T) {
		bank := NewBank()
		bank.AddRate("CHF", "USD", 2)
		result := bank.Reduce(Franc(2), "USD")
		assertMoney(t, Dollar(1), result)
	})

	t.Run("identity rate", func(t *testing.T) {
		bank := NewBank()
		result := bank.Reduce(Franc(1), "CHF")
		assertMoney(t, Franc(1), result)
	})
}
