package money

type Expression interface {
	Reduce(bank Bank, currency string) Money
}

type currencyPair struct {
	from string
	to   string
}

type Bank struct {
	rates map[currencyPair]int
}

func NewBank() Bank {
	rates := make(map[currencyPair]int)
	return Bank{rates: rates}
}

func (bank Bank) rate(from string, to string) int {
	if from == to {
		return 1
	}
	return bank.rates[currencyPair{from: from, to: to}]
}

func (bank Bank) Reduce(source Expression, toCurrency string) Money {
	return source.Reduce(bank, toCurrency)
}

func (bank Bank) AddRate(from string, to string, rate int) {
	bank.rates[currencyPair{from: from, to: to}] = rate
}
