package sort

import (
	"testing"
)

func TestMerge_Sort(t *testing.T) {
	HelperTestSort(t, Merge[int]{})
}

func BenchmarkMerge_Sort(b *testing.B) {
	HelperBenchmarkSort(b, Merge[int]{})
}
