package app

import "rental/internal/domain"

type Repository[T domain.Entity] interface {
	SaveItem(item T) error
	FindByID(id domain.ID) (T, error)
	List() ([]T, error)
}
