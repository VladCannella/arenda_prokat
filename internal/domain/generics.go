package domain

// TODO: Filter[T any], Map[T, U any], Reduce
// TODO: constraint с ~ (например type ID ~string) — зачем нужна тильда

func Filter[T any](items []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(items))
	for _, i := range items {
		if predicate(i) {
			result = append(result, i)
		}
	}
	return result
}

func Map[T, U any](items []T, fn func(T) U) []U {
	result := make([]U, 0, len(items))
	for _, i := range items {
		newItem := fn(i)
		result = append(result, newItem)
	}
	return result
}

func Reduce[T, U any](items []T, initial U, fn func(acc U, items T) U) U {
	acc := initial
	for _, item := range items {
		acc = fn(acc, item)
	}
	return acc
}
