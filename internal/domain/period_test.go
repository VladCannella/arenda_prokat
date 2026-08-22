package domain

import (
	"testing"
	"time"
)

func TestNewPeriod(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		testStart  time.Time
		testEnd    time.Time
		wantAmount Period
		wantErr    bool
	}{
		{name: "start < end", testStart: now, testEnd: now.Add(10 * time.Minute), wantAmount: Period{start: now, end: now.Add(10 * time.Minute)}, wantErr: false},
		{name: "start > end", testStart: now, testEnd: now.Add(-10 * time.Minute), wantAmount: Period{}, wantErr: true},
		{name: "start = end", testStart: now, testEnd: now, wantAmount: Period{}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := NewPeriod(tt.testStart, tt.testEnd)

			if (err != nil) != tt.wantErr {
				t.Errorf("NewPeriod(%v,%v) = %v error = %v, wantErr = %v", tt.testStart, tt.testEnd, got, err, tt.wantErr)
				return
			}

			if !tt.wantErr && got != tt.wantAmount {
				t.Errorf("NewPeriod(%v, %v) = %v, wantAmount = %v", tt.testStart, tt.testEnd, got, tt.wantAmount)
			}

		})
	}
}

func TestDays(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name       string
		testPeriod Period
		wantAmount int64
	}{
		{name: "empty Period", testPeriod: Period{}, wantAmount: 0},
		{name: "Period: 2 days", testPeriod: Period{start: now, end: now.Add(25 * time.Hour)}, wantAmount: 2},
		{name: "Period: 1 day (1 hour entry)", testPeriod: Period{start: now, end: now.Add(2 * time.Hour)}, wantAmount: 1},
		{name: "Period: 3 days", testPeriod: Period{start: now, end: now.Add(49 * time.Hour)}, wantAmount: 3},
	}

	for _, tt := range tests {
		beforePeriod := tt.testPeriod

		t.Run(tt.name, func(t *testing.T) {
			got := tt.testPeriod.Days()

			if got != tt.wantAmount {
				t.Errorf("(%v) Days() = %v, wantAmount: %v", tt.testPeriod, got, tt.wantAmount)
			}

			if beforePeriod != tt.testPeriod {
				t.Errorf("testPeriod mutated after Days(), was %v, now %v", beforePeriod, tt.testPeriod)
			}
		})
	}
}
