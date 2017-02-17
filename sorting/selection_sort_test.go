package sorting

import (
	"testing"

	"github.com/billjh/algo/sorting/utils"
)

func TestSelectionSort(t *testing.T) {
	arr := utils.GetIntArray(10000)
	SelectionSort(arr)
	if !utils.IsSorted(arr) {
		t.Error()
	}
}

func benchmarkSelectionSort(size int, b *testing.B) {
	arr := utils.GetIntArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SelectionSort(arr)
	}
}

func BenchmarkSelectionSort100(b *testing.B)   { benchmarkSelectionSort(100, b) }
func BenchmarkSelectionSort1000(b *testing.B)  { benchmarkSelectionSort(1000, b) }
func BenchmarkSelectionSort10000(b *testing.B) { benchmarkSelectionSort(10000, b) }
