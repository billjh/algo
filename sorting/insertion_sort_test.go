package sorting

import (
	"testing"

	"github.com/billjh/algo/sorting/utils"
)

func TestInsertionSort10000(t *testing.T) {
	arr := utils.GetIntArray(10000)
	InsertionSort(arr)
	if !utils.IsSorted(arr) {
		t.Error("array is not sorted:", arr)
	}
}

func benchmarkInsertionSort(size int, b *testing.B) {
	arr := utils.GetIntArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertionSort(arr)
	}
}

func BenchmarkInsertionSort100(b *testing.B)   { benchmarkInsertionSort(100, b) }
func BenchmarkInsertionSort1000(b *testing.B)  { benchmarkInsertionSort(1000, b) }
func BenchmarkInsertionSort10000(b *testing.B) { benchmarkInsertionSort(10000, b) }
