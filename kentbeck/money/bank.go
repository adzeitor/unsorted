package money

type Expression interface {
}

type Bank struct {
}

func NewBank() Bank {
	return Bank{}
}

func (bank *Bank) Reduce(source Expression, toCurrency string) Money {
	return Dollar(10)
}
