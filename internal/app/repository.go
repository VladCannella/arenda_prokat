package app

import "rental/internal/domain"

type Repository[T domain.Entity] interface {
	Save(item T) error
	FindByID(id domain.ID) (T, error)
	List() ([]T, error)
}
