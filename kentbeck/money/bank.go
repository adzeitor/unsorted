package money

type Expression interface {
	Reduce(currency string) Money
}

type Bank struct {
}

func NewBank() Bank {
	return Bank{}
}

func (bank *Bank) Reduce(source Expression, toCurrency string) Money {
	return source.Reduce(toCurrency)
}
