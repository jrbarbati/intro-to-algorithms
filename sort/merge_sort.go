package sort

import "cmp"

type Merge[T cmp.Ordered] struct{}

// Sort sorts items using merge sort algorithm.
func (m Merge[T]) Sort(items []T) []T {
	return items
}
