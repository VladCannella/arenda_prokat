package domain

import "testing"

func TestDeferArgEval(t *testing.T) {
	captured, final := DeferArgEval()
	if captured != 1 {
		t.Errorf("captured = %d, want 1", captured)
	}
	if final != 3 {
		t.Errorf("final = %d, want 3", final)
	}
}
