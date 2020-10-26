package money

type DollarValue struct {
	amount int
}

func (dollar DollarValue) Times(n int) DollarValue {
	return Dollar(dollar.amount * n)
}

func (dollar DollarValue) Equal(other interface{}) bool {
	otherDollar := other.(DollarValue)
	return dollar.amount == otherDollar.amount
}

func Dollar(amount int) DollarValue {
	return DollarValue{amount: amount}
}
