package money

type Money struct {
	amount int
}

func (money Money) Equal(other Money) bool {
	return money.amount == other.amount
}

type DollarValue struct {
	Money
}

func (dollar DollarValue) Times(n int) DollarValue {
	return Dollar(dollar.amount * n)
}

func (dollar DollarValue) Equal(other DollarValue) bool {
	return dollar.Money.Equal(other.Money)
}

func Dollar(amount int) DollarValue {
	return DollarValue{
		Money: Money{
			amount: amount,
		},
	}
}
