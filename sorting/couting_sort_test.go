package sorting

import (
	"testing"

	"github.com/billjh/algo/sorting/utils"
)

func TestCountingSort10000(t *testing.T) {
	k := 20000
	arr := utils.GetIntNArray(10000, k)
	CountingSort(arr, k)
	if !utils.IsSorted(arr) {
		t.Error(arr)
	}
}

func benchmarkCountingSort(size int, b *testing.B) {
	k := 20000
	arr := utils.GetIntNArray(size, k)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountingSort(arr, k)
	}
}

func BenchmarkCountingSort100(b *testing.B)   { benchmarkCountingSort(100, b) }
func BenchmarkCountingSort1000(b *testing.B)  { benchmarkCountingSort(1000, b) }
func BenchmarkCountingSort10000(b *testing.B) { benchmarkCountingSort(10000, b) }
