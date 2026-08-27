package domain

import (
	"errors"
	"fmt"
)

var ErrItemAlreadyRented = errors.New("item: item already rented")
var ErrItemNotFound = errors.New("item: item not found")
var ErrRentalClosed = errors.New("rental: rental is closed")

type ValidationError struct {
	Field  string
	Reason string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("field %q failed the test: %q", e.Field, e.Reason)
}

type DomainError struct {
	Op  string
	Err error
}

func (e DomainError) Error() string {
	return fmt.Sprintf("%v: %v", e.Op, e.Err)
}

func (e DomainError) Unwrap() error {
	return e.Err
}
