package domain

import (
	"errors"
	"testing"
)

func TestValidationError(t *testing.T) {
	tests := []struct {
		name      string
		errorTest ValidationError
		wantError string
	}{
		{name: "empty amount", errorTest: ValidationError{Field: "amount", Reason: "must not be empty"}, wantError: "field \"amount\" failed the test: \"must not be empty\""},
		{name: "one value period", errorTest: ValidationError{Field: "period", Reason: "must be two values"}, wantError: "field \"period\" failed the test: \"must be two values\""},
		{name: "empty name", errorTest: ValidationError{Field: "customer", Reason: "name is required field"}, wantError: "field \"customer\" failed the test: \"name is required field\""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.errorTest.Error() != tt.wantError {
				t.Errorf("(%q) Error() = %q, wanteError = %q", tt.errorTest, tt.errorTest.Error(), tt.wantError)
			}
		})
	}
}

func TestDomainError_Error(t *testing.T) {
	tests := []struct {
		name      string
		errorTest DomainError
		wantError string
	}{
		{name: "item already rented", errorTest: DomainError{Op: "rent item item-42", Err: ErrItemAlreadyRented}, wantError: "rent item item-42: item: item already rented"},
		{name: "item not found", errorTest: DomainError{Op: "rent item item-42", Err: ErrItemNotFound}, wantError: "rent item item-42: item: item not found"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.errorTest.Error() != tt.wantError {
				t.Errorf("(%q) Error() = %q, wanteError = %q", tt.errorTest, tt.errorTest.Error(), tt.wantError)
			}
		})
	}
}

func TestDomainError_Unwrap_ErrorsIs(t *testing.T) {
	tests := []struct {
		name      string
		errorTest DomainError
		checkErr  error
		wantErr   bool
	}{
		{name: "item already rented", errorTest: DomainError{Op: "rent item item-42", Err: ErrItemAlreadyRented}, checkErr: ErrItemAlreadyRented, wantErr: true},
		{name: "item not found", errorTest: DomainError{Op: "rent item item-42", Err: ErrItemNotFound}, checkErr: ErrItemAlreadyRented, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errChecker := errors.Is(tt.errorTest, tt.checkErr)
			if errChecker != tt.wantErr {
				t.Errorf("errors.Is(errorTest, %v) = %v, want %v", tt.checkErr, errChecker, tt.wantErr)
			}
		})
	}
}

func TestErrorAs_ValidationError(t *testing.T) {
	tests := []struct {
		name       string
		errorTest  error
		wantFound  bool
		wantField  string
		wantReason string
	}{
		{name: "item must not be empty", errorTest: DomainError{Op: "rent item item-42", Err: ValidationError{Field: "item", Reason: "must not be empty"}}, wantFound: true, wantField: "item", wantReason: "must not be empty"},
		{name: "item is unavailable", errorTest: DomainError{Op: "rent item item-43", Err: ValidationError{Field: "item", Reason: "item is not available"}}, wantFound: true, wantField: "item", wantReason: "item is not available"},
		{name: "without ValidationError", errorTest: DomainError{Op: "rent item item-44", Err: ErrCurrencyMismatch}, wantFound: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var found ValidationError

			ok := errors.As(tt.errorTest, &found)

			if ok != tt.wantFound {
				t.Errorf("error.As found = %v, want = %v", found, ValidationError{Field: tt.wantField, Reason: tt.wantReason})
				return
			}
			if tt.wantFound && (found.Reason != tt.wantReason || found.Field != tt.wantField) {
				t.Errorf("wantField:%v, get: %v. wantReason:%v, get: %v", tt.wantField, found.Field, tt.wantReason, found.Reason)
			}
		})
	}
}
