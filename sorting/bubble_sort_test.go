package sorting

import (
	"testing"

	"github.com/billjh/algo/sorting/utils"
)

func TestBubbleSort(t *testing.T) {
	arr := utils.GetIntArray(10000)
	BubbleSort(arr)
	if !utils.IsSorted(arr) {
		t.Error()
	}
}

func benchmarkBubbleSort(size int, b *testing.B) {
	arr := utils.GetIntArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BubbleSort(arr)
	}
}

func BenchmarkBubbleSort100(b *testing.B)   { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort1000(b *testing.B)  { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B) { benchmarkBubbleSort(10000, b) }
