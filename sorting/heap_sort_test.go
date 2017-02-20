package sorting

import "testing"

import "github.com/billjh/algo/sorting/utils"

func TestHeapify(t *testing.T) {
	arr := utils.GetIntArray(10000)
	heapify(arr)
	if !utils.IsMaxHeap(arr) {
		t.Error()
	}
}

func TestSiftDown(t *testing.T) {
	testCases := [][]int{
		[]int{3, 2, 1},
		[]int{3, 4, 7, 3, 2},
		[]int{3, 8, 9, 5, 4},
	}
	for _, testCase := range testCases {
		siftDown(testCase, 0, len(testCase))
		if !utils.IsMaxHeap(testCase) {
			t.Error(testCase, "not heapified")
		}
	}
}

func TestHeapSort10000(t *testing.T) {
	arr := utils.GetIntArray(10000)
	HeapSort(arr)
	if !utils.IsSorted(arr) {
		t.Error()
	}
}

func benchmarkHeapSort(size int, b *testing.B) {
	arr := utils.GetIntArray(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		HeapSort(arr)
	}
}

func BenchmarkHeapSort100(b *testing.B)   { benchmarkHeapSort(100, b) }
func BenchmarkHeapSort1000(b *testing.B)  { benchmarkHeapSort(1000, b) }
func BenchmarkHeapSort10000(b *testing.B) { benchmarkHeapSort(10000, b) }
