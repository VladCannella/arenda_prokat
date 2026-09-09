package infra

import (
	"errors"
	"rental/internal/app"
	"rental/internal/domain"
	"sync"
)

var ErrEmptyID = errors.New("repository: ID is empty")

type InMemoryRepo[T domain.Entity] struct {
	mu   sync.Mutex
	data map[domain.ID]T
}

var _ app.Repository[domain.Item] = (*InMemoryRepo[domain.Item])(nil)

func NewInMemoryRepo[T domain.Entity]() *InMemoryRepo[T] {
	return &InMemoryRepo[T]{data: make(map[domain.ID]T)}
}

func (r *InMemoryRepo[T]) Save(item T) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if item.GetID() != "" {
		r.data[item.GetID()] = item
		return nil
	}
	return ErrEmptyID

}

func (r *InMemoryRepo[T]) FindByID(id domain.ID) (T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, ok := r.data[id]
	if !ok {
		var zero T
		return zero, domain.ErrEntityNotFound
	}
	return item, nil
}

func (r *InMemoryRepo[T]) List() ([]T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	list := make([]T, 0, len(r.data))
	for _, v := range r.data {
		list = append(list, v)
	}
	return list, nil
}
