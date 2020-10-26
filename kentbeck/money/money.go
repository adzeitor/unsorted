package money

type DollarValue struct {
	Amount int
}

func (dollar *DollarValue) Times(n int) {
	dollar.Amount *= n
}

func Dollar(amount int) DollarValue {
	return DollarValue{Amount: amount}
}
