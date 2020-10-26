package money

import (
	"testing"
)

func TestBank_Reduce(t *testing.T) {
	t.Run("reduce sum", func(t *testing.T) {
		sum := Dollar(3).Plus(Dollar(4))
		bank := Bank{}

		reduced := bank.Reduce(sum, "USD")
		assertMoney(t, Dollar(7), reduced)
	})

	t.Run("reduce money", func(t *testing.T) {
		bank := Bank{}
		reduced := bank.Reduce(Dollar(1), "USD")
		assertMoney(t, Dollar(1), reduced)
	})
}
