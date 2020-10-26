package money

type Money struct {
	amount   int
	currency string
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

type DollarValue struct {
	Money
}

func (dollar DollarValue) Times(n int) DollarValue {
	return DollarValue{Money: dollar.Money.Times(n)}
}

func (dollar DollarValue) Equal(other DollarValue) bool {
	return dollar.Money.Equal(other.Money)
}

func Dollar(amount int) DollarValue {
	return DollarValue{
		Money: Money{
			amount:   amount,
			currency: "USD",
		},
	}
}
