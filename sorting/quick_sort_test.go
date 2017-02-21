package sorting

import (
	"testing"

	"github.com/billjh/algo/sorting/utils"
)

func TestPartition(t *testing.T) {
	arr := utils.GetIntArray(10000)
	p := partition(arr, 0, len(arr)-1)
	pivot := arr[p]
	for i, n := range arr {
		if i < p && n > pivot {
			t.Error()
		}
		if i > p && n < pivot {
			t.Error()
		}
	}
}

func TestQuickSort10000(t *testing.T) {
	arr := utils.GetIntArray(10000)
	QuickSort(arr)
	if !utils.IsSorted(arr) {
		t.Error()
	}
}

func benchmarkQuickSort(size int, b *testing.B) {
	arr := utils.GetIntArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSort(arr)
	}
}

func BenchmarkQuickSort100(b *testing.B)   { benchmarkQuickSort(100, b) }
func BenchmarkQuickSort1000(b *testing.B)  { benchmarkQuickSort(1000, b) }
func BenchmarkQuickSort10000(b *testing.B) { benchmarkQuickSort(10000, b) }
