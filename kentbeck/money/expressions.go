package money

type Sum struct {
	augend Expression
	addend Expression
}

func (sum Sum) Reduce(bank Bank, currency string) Money {
	augend := sum.augend.Reduce(bank, currency)
	addend := sum.addend.Reduce(bank, currency)
	return Money{
		amount:   augend.amount + addend.amount,
		currency: currency,
	}
}

func (sum Sum) Plus(addend Expression) Expression {
	return nil
}
