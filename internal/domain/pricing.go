package domain

import "errors"

// TODO: PricingStrategy — интерфейс с одним методом (OCP)
// TODO: реализации DailyPricing, WeeklyDiscountPricing, LoyalCustomerPricing —
// решить, в каком слое им место (интерфейс объявляется там, где используется)

var ErrInvalidDiscount = errors.New("pricing: discount value is invalid")

type PricingStrategy interface {
	Calculate(period Period, dailyRate Money, customer Customer) (Money, error)
}

// DailyPricing
type DailyPricing struct{}

func (p DailyPricing) Calculate(period Period, dailyRate Money, customer Customer) (Money, error) {
	return dailyRate.MulInt(period.Days()), nil
}

// WeeklyDiscount
func NewWeeklyDiscountPricing(days int64, salePercents int64) (WeeklyDiscountPricing, error) {
	if salePercents > 0 && salePercents < 101 {
		return WeeklyDiscountPricing{limitDays: days, sale: salePercents}, nil
	}
	return WeeklyDiscountPricing{}, ErrInvalidDiscount
}

type WeeklyDiscountPricing struct {
	limitDays int64
	sale      int64
}

func (p WeeklyDiscountPricing) Calculate(period Period, dailyRate Money, customer Customer) (Money, error) {
	pricing := dailyRate.MulInt(period.Days())
	if period.Days() >= p.limitDays {
		pricing.amountMinor = pricing.amountMinor * (100 - p.sale) / 100
	}
	return pricing, nil
}

// LoyalCustomerPricing

type LoyalCustomerPricing struct{}

func (p LoyalCustomerPricing) Calculate(period Period, dailyRate Money, customer Customer) (Money, error) {
	var fixsale int64 = 70
	pricing := dailyRate.MulInt(period.Days())
	if customer.RentalCount > 100 {
		pricing.amountMinor = pricing.amountMinor * (100 - fixsale) / 100
	}
	return pricing, nil
}
