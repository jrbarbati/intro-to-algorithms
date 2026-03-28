package sort

import "cmp"

type Bubble[T cmp.Ordered] struct{}

func (Bubble[T]) Sort(items []T) []T {
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			if items[i] < items[j] {
				continue
			}

			items[i], items[j] = items[j], items[i]
		}
	}

	return items
}
