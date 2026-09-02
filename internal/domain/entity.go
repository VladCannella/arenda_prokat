package domain

import "time"

type ID string

type BaseEntity struct {
	ID        ID
	CreatedAt time.Time
}
