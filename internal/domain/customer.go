package domain

type Customer struct {
	BaseEntity
	Name        string
	RentalCount int64
}
