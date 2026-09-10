package domain

func DeferArgEval() (capturedAtDefer int, finalValue int) {
	x := 1
	defer func(v int) {
		capturedAtDefer = v
	}(x)

	x = 2
	x = 3
	finalValue = x
	return
}
