package sort

import (
	"testing"
)

func TestSelection_Sort(t *testing.T) {
	HelperTestSort(t, Selection[int]{})
}

func BenchmarkSelection_Sort(b *testing.B) {
	HelperBenchmarkSort(b, Selection[int]{})
}
