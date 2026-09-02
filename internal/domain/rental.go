package domain

import "time"

type Rental struct {
	BaseEntity
	CustomerID     ID
	ItemID         ID
	Period         Period
	ActualReturnAt time.Time
}
