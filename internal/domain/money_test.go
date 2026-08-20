package domain

import (
	"testing"
)

func TestNewMoney(t *testing.T) {
	tests := []struct {
		name         string
		amountTest   int64
		currencyTest string
		wantAmount   Money
		wantErr      bool
	}{
		{name: "EUR", amountTest: 100, currencyTest: "EUR", wantAmount: Money{100, "EUR"}, wantErr: false},
		{name: "USD", amountTest: 100, currencyTest: "USD", wantAmount: Money{100, "USD"}, wantErr: false},
		{name: "Empty Currency", amountTest: 100, currencyTest: "", wantAmount: Money{}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMoney(tt.amountTest, tt.currencyTest)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMoney(%d, %q) error = %v, wantErr = %v", tt.amountTest, tt.currencyTest, err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.wantAmount {
				t.Errorf("NewMoney(%d, %q) = %v, want %v", tt.amountTest, tt.currencyTest, got, tt.wantAmount)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name       string
		a, b       Money
		wantAmount Money
		wantErr    bool
	}{
		{name: "RUB + RUB", a: Money{100, "RUB"}, b: Money{100, "RUB"}, wantAmount: Money{200, "RUB"}, wantErr: false},
		{name: "EUR + EUR", a: Money{150, "EUR"}, b: Money{150, "EUR"}, wantAmount: Money{300, "EUR"}, wantErr: false},
		{name: "RUB + EUR", a: Money{200, "RUB"}, b: Money{100, "EUR"}, wantAmount: Money{}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beforeA := tt.a
			beforeB := tt.b

			got, err := tt.a.Add(tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("(%+v) Add(%+v) error = %v, wantErr = %v", tt.a, tt.b, err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.wantAmount {
				t.Errorf("(%+v) Add(%+v) = %+v, want %v", tt.a, tt.b, got, tt.wantAmount)
			}

			if tt.a != beforeA {
				t.Errorf("a mutated after Add: was %v, now: %v", beforeA, tt.a)
			}

			if tt.b != beforeB {
				t.Errorf("b mutated after Add: was %v, now %v", beforeB, tt.b)
			}
		})
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		name       string
		a, b       Money
		wantAmount Money
		wantErr    bool
	}{
		{name: "EUR - EUR", a: Money{300, "EUR"}, b: Money{222, "EUR"}, wantAmount: Money{78, "EUR"}, wantErr: false},
		{name: "USD - USD", a: Money{200, "USD"}, b: Money{50, "USD"}, wantAmount: Money{150, "USD"}, wantErr: false},
		{name: "EUR - USD", a: Money{222, "EUR"}, b: Money{222, "USD"}, wantAmount: Money{}, wantErr: true},
		{name: "USD - EUR", a: Money{456, "USD"}, b: Money{222, "EUR"}, wantAmount: Money{}, wantErr: true},
		{name: "Negative value", a: Money{50, "USD"}, b: Money{100, "USD"}, wantAmount: Money{-50, "USD"}, wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beforeA := tt.a
			beforeB := tt.b

			got, err := tt.a.Sub(tt.b)

			if (err != nil) != tt.wantErr {
				t.Errorf("(%+v) Sub (%+v) error = %v, wantErr = %v", tt.a, tt.b, err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.wantAmount {
				t.Errorf("(%+v) Sub(%+v) = %+v, want: %+v", tt.a, tt.b, got, tt.wantAmount)
			}

			if tt.a != beforeA {
				t.Errorf("a mutated after Sub: was %v, now %v", beforeA, tt.a)
			}
			if tt.b != beforeB {
				t.Errorf("b mutated after Sub: was %v, now %v", beforeB, tt.b)
			}
		})
	}
}

func TestMulInt(t *testing.T) {
	tests := []struct {
		name       string
		a          Money
		number     int64
		wantAmount Money
	}{
		{name: "mul 0", a: Money{100, "USD"}, number: 0, wantAmount: Money{0, "USD"}},
		{name: "mul 10", a: Money{20, "EUR"}, number: 10, wantAmount: Money{200, "EUR"}},
		{name: "mul 1", a: Money{50, "USD"}, number: 1, wantAmount: Money{50, "USD"}},
		{name: "mul -1", a: Money{300, "RUB"}, number: -1, wantAmount: Money{-300, "RUB"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beforeA := tt.a

			got := tt.a.MulInt(tt.number)

			if got != tt.wantAmount {
				t.Errorf("(%v) MulInt(%v) = %v, want %v", tt.a, tt.number, got, tt.wantAmount)
			}

			if tt.a != beforeA {
				t.Errorf("a mutated after MulInt: was %v, now %v", beforeA, tt.a)
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name       string
		a          Money
		wantAmount string
	}{
		{name: "zero value", a: Money{0, "RUB"}, wantAmount: "0.00 RUB"},
		{name: "negative value", a: Money{-10050, "RUB"}, wantAmount: "-100.50 RUB"},
		{name: "positive value", a: Money{10050, "EUR"}, wantAmount: "100.50 EUR"},
		{name: "single-digit cents", a: Money{10005, "USD"}, wantAmount: "100.05 USD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beforeA := tt.a

			got := tt.a.String()

			if got != tt.wantAmount {
				t.Errorf("(%v) String() = %v, want %v", tt.a, got, tt.wantAmount)
			}

			if beforeA != tt.a {
				t.Errorf("a mutated after String(): was %v, now %v", beforeA, tt.a)
			}
		})
	}
}

func TestEquality(t *testing.T) {
	t.Run("TestEqual: amount + currency", func(t *testing.T) {
		a := Money{100, "RUB"}
		b := Money{100, "RUB"}
		if a != b {
			t.Errorf("a and b is not equal: a %v, b %v", a, b)
		}
	})
	t.Run("TestEqual: currency", func(t *testing.T) {
		a := Money{100, "RUB"}
		b := Money{100, "USD"}
		if a == b {
			t.Errorf("a and b is equal: a %v, b %v", a, b)
		}
	})
	t.Run("TestEqual: amount", func(t *testing.T) {
		a := Money{200, "RUB"}
		b := Money{100, "RUB"}
		if a == b {
			t.Errorf("a and b is equal: a %v, b %v", a, b)
		}
	})
}
