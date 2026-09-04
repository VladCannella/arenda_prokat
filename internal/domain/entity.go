package domain

import "time"

type ID string

type Entity interface {
	GetID() ID
}

type BaseEntity struct {
	ID        ID
	CreatedAt time.Time
}

func (e BaseEntity) GetID() ID {
	return e.ID
}
