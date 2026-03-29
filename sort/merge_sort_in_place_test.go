package sort

import (
	"testing"
)

func TestMergeInPlace_Sort(t *testing.T) {
	HelperTestSort(t, MergeInPlace[int]{})
}

func BenchmarkMergeInPlace_Sort(b *testing.B) {
	HelperBenchmarkSort(b, MergeInPlace[int]{})
}
