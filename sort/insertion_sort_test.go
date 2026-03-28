package sort

import (
	"testing"
)

func TestInsertion_Sort(t *testing.T) {
	HelperTestSort(t, Insertion[int]{})
}

func BenchmarkInsertion_Sort(b *testing.B) {
	HelperBenchmarkSort(b, Insertion[int]{})
}
