package money

type FrancValue struct {
	Money
}

func (franc FrancValue) Times(n int) FrancValue {
	return Franc(franc.amount * n)
}

func (franc FrancValue) Equal(other FrancValue) bool {
	return franc.Money.Equal(other.Money)
}

func Franc(amount int) FrancValue {
	return FrancValue{
		Money: Money{
			amount:   amount,
			currency: "CHF",
		},
	}
}
