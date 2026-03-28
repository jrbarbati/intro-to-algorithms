package sort

import (
	"testing"
)

func TestBubble_Sort(t *testing.T) {
	HelperTestSort(t, Bubble[int]{})
}

func BenchmarkBubble_Sort(b *testing.B) {
	HelperBenchmarkSort(b, Bubble[int]{})
}
