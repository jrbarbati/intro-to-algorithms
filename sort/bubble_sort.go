package sort

import "cmp"

type Bubble[T cmp.Ordered] struct{}

// Sort sorts items using bubble sort algorithm.
func (Bubble[T]) Sort(items []T) []T {
	for i := 0; i < len(items); i++ {
		for j := 0; j < len(items)-(i+1); j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}

	return items
}
