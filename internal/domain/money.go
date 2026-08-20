package domain

import (
	"errors"
	"fmt"
)

// TODO: Money — неизменяемый value object (int64 минорные единицы + валюта)
// TODO: Add, Sub, MulInt, String() — без мутации получателя
// TODO: сложение разных валют — ошибка, а не паника
// TODO: Money должен быть comparable (a == b) — подумать про поля структуры
type Money struct {
	amountMinor int64
	currency    string
}

var ErrEmptyCurrency = errors.New("money: currency is required")
var ErrCurrencyMismatch = errors.New("money: currency mismatch")

func NewMoney(amountMinor int64, currency string) (Money, error) {
	if currency == "" {
		return Money{}, ErrEmptyCurrency
	}
	return Money{amountMinor: amountMinor, currency: currency}, nil
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency == other.currency {
		newAmount := m.amountMinor + other.amountMinor
		return Money{amountMinor: newAmount, currency: m.currency}, nil
	}

	return Money{}, ErrCurrencyMismatch
}

func (m Money) Sub(other Money) (Money, error) {
	if m.currency == other.currency {
		newAmount := m.amountMinor - other.amountMinor
		return Money{amountMinor: newAmount, currency: m.currency}, nil
	}

	return Money{}, ErrCurrencyMismatch
}

func (m Money) MulInt(factor int64) Money {
	return Money{amountMinor: m.amountMinor * factor, currency: m.currency}
}

func (m Money) String() string {
	sign := ""
	if m.amountMinor < 0 {
		m.amountMinor = m.amountMinor * (-1)
		sign = "-"
	}

	newString := fmt.Sprintf("%s%d.%02d %s", sign, m.amountMinor/100, m.amountMinor%100, m.currency)
	return newString
}
