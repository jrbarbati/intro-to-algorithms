package sort

import "cmp"

type Quick[T cmp.Ordered] struct{}

// Sort sorts items using quick sort algorithm.
func (q Quick[T]) Sort(items []T) []T {
	return items
}
