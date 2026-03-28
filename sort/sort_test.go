package sort

import (
	"math/rand"
	"testing"
)

func HelperTestSort(t *testing.T, s Sorter[int]) {
	scenarios := []struct {
		name     string
		actual   []int
		expected []int
	}{
		{
			name:     "nil",
			actual:   nil,
			expected: nil,
		},
		{
			name:     "empty",
			actual:   []int{},
			expected: []int{},
		},
		{
			name:     "unsorted",
			actual:   []int{92, 384, 1331, 1038, 1, 231, 321, 49432, 948},
			expected: []int{1, 92, 231, 321, 384, 948, 1038, 1331, 49432},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := s.Sort(scenario.actual)

			for i, actual := range result {
				if actual != scenario.expected[i] {
					t.Errorf("actual %v != expected %v", actual, scenario.expected[i])
				}
			}
		})
	}
}

func HelperBenchmarkSort(b *testing.B, s Sorter[int]) {
	arr := make([]int, 1_000)

	for i := 0; i < 1_000; i++ {
		arr[i] = rand.Intn(1_000_000)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		s.Sort(arr)
	}
}
