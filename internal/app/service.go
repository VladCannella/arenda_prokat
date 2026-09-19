package app

import (
	"fmt"
	"rental/internal/domain"
	"time"
)

type RentalService struct {
	items     Repository[domain.Item]
	customers Repository[domain.Customer]
	rentals   Repository[domain.Rental]
	notify    Notifier
	price     domain.PricingStrategy
}

func NewRentalService(
	item Repository[domain.Item],
	customer Repository[domain.Customer],
	rental Repository[domain.Rental],
	notify Notifier,
	price domain.PricingStrategy,
) *RentalService {
	return &RentalService{
		items:     item,
		customers: customer,
		rentals:   rental,
		notify:    notify,
		price:     price,
	}
}

func (s *RentalService) RentItem(itemID, customerID domain.ID, start, end time.Time) (rental domain.Rental, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("rent item %s: %w", itemID, err)
		}
	}()
	item, err := s.items.FindByID(itemID)
	if err != nil {
		return domain.Rental{}, err
	}

	if item.Status == domain.ItemRented {
		return domain.Rental{}, domain.ErrItemAlreadyRented
	}
	customer, err := s.customers.FindByID(customerID)
	if err != nil {
		return domain.Rental{}, err
	}

	period, err := domain.NewPeriod(start, end)
	if err != nil {
		return domain.Rental{}, err
	}

	cost, err := s.price.Calculate(period, item.DailyRate, customer)
	if err != nil {
		return domain.Rental{}, err
	}

	rental = domain.Rental{
		BaseEntity: domain.BaseEntity{
			ID:        domain.ID(fmt.Sprintf("rental-%d", time.Now().UnixNano())),
			CreatedAt: time.Now(),
		},
		CustomerID: customerID,
		ItemID:     itemID,
		Period:     period,
	}

	item.Status = domain.ItemRented
	if err := s.items.Save(item); err != nil {
		return domain.Rental{}, err
	}

	if err := s.rentals.Save(rental); err != nil {
		return domain.Rental{}, err
	}

	if err := s.notify.Notify(fmt.Sprintf("вещь %s арендована, стоимость %s", itemID, cost.String())); err != nil {
		return domain.Rental{}, err
	}

	return rental, nil
}
