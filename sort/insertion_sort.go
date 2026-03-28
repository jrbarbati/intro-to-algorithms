package sort

import "cmp"

type Insertion[T cmp.Ordered] struct{}

// Sort sorts items using insertion sort algorithm.
func (Insertion[T]) Sort(items []T) []T {
	for i := 1; i < len(items); i++ {
		for j := i - 1; j >= 0; j-- {
			if items[j+1] >= items[j] {
				break
			}

			items[j], items[j+1] = items[j+1], items[j]
		}
	}

	return items
}
