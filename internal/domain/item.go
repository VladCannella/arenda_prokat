package domain

type ItemStatus string

const (
	ItemAvailable ItemStatus = "available"
	ItemRented    ItemStatus = "rented"
)

type Item struct {
	BaseEntity
	Name      string
	DailyRate Money
	Status    ItemStatus
}
