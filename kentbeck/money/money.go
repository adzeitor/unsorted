package money

type DollarValue struct {
	Amount int
}

func (dollar DollarValue) Times(n int) DollarValue {
	return Dollar(dollar.Amount * n)
}

func (dollar DollarValue) Equal(other interface{}) bool {
	otherDollar := other.(DollarValue)
	return dollar.Amount == otherDollar.Amount
}

func Dollar(amount int) DollarValue {
	return DollarValue{Amount: amount}
}
