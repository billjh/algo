package sorting

import (
	"testing"

	"github.com/billjh/algo/sorting/utils"
)

func TestOptimizedInsertionSort10000(t *testing.T) {
	arr := utils.GetIntArray(10000)
	OptimizedInsertionSort(arr)
	if !utils.IsSorted(arr) {
		t.Error("array is not sorted:", arr)
	}
}

func benchmarkOptimizedInsertionSort(size int, b *testing.B) {
	arr := utils.GetIntArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		OptimizedInsertionSort(arr)
	}
}

func BenchmarkOptimizedInsertionSort100(b *testing.B)   { benchmarkOptimizedInsertionSort(100, b) }
func BenchmarkOptimizedInsertionSort1000(b *testing.B)  { benchmarkOptimizedInsertionSort(1000, b) }
func BenchmarkOptimizedInsertionSort10000(b *testing.B) { benchmarkOptimizedInsertionSort(10000, b) }
