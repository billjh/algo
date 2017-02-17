package sorting

import (
	"testing"

	"github.com/billjh/algo/sorting/utils"
)

func TestMergeSort(t *testing.T) {
	arr := utils.GetIntArray(10000)
	MergeSort(arr)
	if !utils.IsSorted(arr) {
		t.Error()
	}
}

func benchmarkMergeSort(size int, b *testing.B) {
	arr := utils.GetIntArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSort(arr)
	}
}

func BenchmarkMergeSort100(b *testing.B)   { benchmarkMergeSort(100, b) }
func BenchmarkMergeSort1000(b *testing.B)  { benchmarkMergeSort(1000, b) }
func BenchmarkMergeSort10000(b *testing.B) { benchmarkMergeSort(10000, b) }
