package domain

import (
	"errors"
	"fmt"
)

// Money is an immutable object of value representing a monetary amount
// in insignificant units (e.g., cents, kopecks). Comparable with ==.
type Money struct {
	amountMinor int64
	currency    string
}

// ErrEmptyCurrency appears when the currency name field is empty
var ErrEmptyCurrency = errors.New("money: currency is required")

// ErrCurrencyMismatch appears when different currencies are specified in currency transactions (we only work with the same ones)
var ErrCurrencyMismatch = errors.New("money: currency mismatch")

// NewMoney validates the input data;
// If the currency is empty, return the ErrEmptyCurrency error.
func NewMoney(amountMinor int64, currency string) (Money, error) {
	if currency == "" {
		return Money{}, ErrEmptyCurrency
	}
	return Money{amountMinor: amountMinor, currency: currency}, nil
}

// Add function returns the sum of m and other.
// Returns ErrCurrencyMismatch if their currencies differ.
func (m Money) Add(other Money) (Money, error) {
	if m.currency == other.currency {
		newAmount := m.amountMinor + other.amountMinor
		return Money{amountMinor: newAmount, currency: m.currency}, nil
	}

	return Money{}, ErrCurrencyMismatch
}

// Sub function returns the subsctract between m and other.
// Returns ErrCurrencyMismatch if their currencies differ.
func (m Money) Sub(other Money) (Money, error) {
	if m.currency == other.currency {
		newAmount := m.amountMinor - other.amountMinor
		return Money{amountMinor: newAmount, currency: m.currency}, nil
	}

	return Money{}, ErrCurrencyMismatch
}

// MulInt multiplies the m.amountMinor by an arbitrary number
func (m Money) MulInt(factor int64) Money {
	return Money{amountMinor: m.amountMinor * factor, currency: m.currency}
}

// String converts a number to a string to output a valid value of type "<amount> <currency >"
// Negative numbers in the entry are processed modulo, and then "-" is added to the string.
func (m Money) String() string {
	sign := ""
	if m.amountMinor < 0 {
		m.amountMinor = m.amountMinor * (-1)
		sign = "-"
	}

	newString := fmt.Sprintf("%s%d.%02d %s", sign, m.amountMinor/100, m.amountMinor%100, m.currency)
	return newString
}
