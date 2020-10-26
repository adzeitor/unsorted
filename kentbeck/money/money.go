package money

import "fmt"

type Money struct {
	amount   int
	currency string
}

func (money Money) String() string {
	return fmt.Sprintf("%d %s", money.amount, money.currency)
}

func (money Money) Equal(other Money) bool {
	return money.amount == other.amount &&
		money.currency == other.currency
}

func (money Money) Times(n int) Money {
	return Money{
		amount:   money.amount * n,
		currency: money.currency,
	}
}

func (money Money) Plus(other Expression) Expression {
	return Sum{
		augend: money,
		addend: other,
	}
}

func (money Money) Reduce(bank Bank, toCurrency string) Money {
	rate := bank.rate(money.currency, toCurrency)
	return Money{
		amount:   money.amount / rate,
		currency: toCurrency,
	}
}

func Dollar(amount int) Money {
	return Money{
		amount:   amount,
		currency: "USD",
	}
}

func Franc(amount int) Money {
	return Money{
		amount:   amount,
		currency: "CHF",
	}
}
