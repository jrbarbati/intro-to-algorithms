package sort

import (
	"testing"
)

func TestQuick_Sort(t *testing.T) {
	HelperTestSort(t, Quick[int]{})
}

func BenchmarkQuick_Sort(b *testing.B) {
	HelperBenchmarkSort(b, Quick[int]{})
}
