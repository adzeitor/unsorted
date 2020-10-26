package money

type Sum struct {
	augend Money
	addend Money
}

func (sum Sum) Reduce(currency string) Money {
	return Money{
		amount:   sum.augend.amount + sum.addend.amount,
		currency: currency,
	}
}
