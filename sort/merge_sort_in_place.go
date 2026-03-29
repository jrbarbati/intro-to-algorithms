package sort

import (
	"cmp"
)

type MergeInPlace[T cmp.Ordered] struct{}

// Sort sorts items using merge sort algorithm.
func (m MergeInPlace[T]) Sort(items []T) []T {
	return m.sort(items)
}

func (m MergeInPlace[T]) sort(items []T) []T {
	if len(items) <= 1 {
		return items
	}

	mid := len(items) / 2

	left := m.sort(items[:mid])
	right := m.sort(items[mid:])

	return m.merge(left, right)
}

func (m MergeInPlace[T]) merge(left, right []T) []T {
	// TODO: Refactor to merge in-place, rather than creating new arrays
	var i, j, k int
	merged := make([]T, 0, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}

		k++
	}

	if i < len(left) {
		merged = append(merged, left[i:]...)
	}

	if j < len(right) {
		merged = append(merged, right[j:]...)
	}

	return merged
}
