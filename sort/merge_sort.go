package sort

import "cmp"

type Merge[T cmp.Ordered] struct{}

func (Merge[T]) Sort(items []T) []T {
	return items
}
