package money

type FrancValue struct {
	amount int
}

func (franc FrancValue) Times(n int) FrancValue {
	return Franc(franc.amount * n)
}

func (franc FrancValue) Equal(other interface{}) bool {
	otherFranc := other.(FrancValue)
	return franc.amount == otherFranc.amount
}

func Franc(amount int) FrancValue {
	return FrancValue{amount: amount}
}
