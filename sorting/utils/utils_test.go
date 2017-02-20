package utils

import "testing"

func TestGetIntArray(t *testing.T) {
	testCases := []int{
		0,
		1,
		666,
	}
	for _, testCase := range testCases {
		arr := GetIntArray(testCase)
		if len(arr) != testCase {
			t.Error("array length not match:", testCase, arr)
		}
	}
}

func TestIsSortedTrue(t *testing.T) {
	var testCases = [][]int{
		[]int{},
		[]int{1},
		[]int{1, 3, 5},
		[]int{6, 6, 6},
		[]int{318, 1363, 1480, 1812, 4132, 6505, 7487, 8129, 9199, 9729},
	}
	for _, testCase := range testCases {
		if !IsSorted(testCase) {
			t.Error(testCase, "should return true")
		}
	}
}

func TestIsSortedFalse(t *testing.T) {
	testCases := [][]int{
		[]int{10, 0},
		[]int{9, 6, 3},
	}
	for _, testCase := range testCases {
		if IsSorted(testCase) {
			t.Error(testCase, "should return false")
		}
	}
}

func TestIsMaxHeapTrue(t *testing.T) {
	var maxHeaps = [][]int{
		[]int{1},
		[]int{6, 1, 3},
		[]int{7, 5, 2, 4, 3, 2, 1},
	}
	for _, heap := range maxHeaps {
		if !IsMaxHeap(heap) {
			t.Error(heap, "should return true")
		}
	}
}

func TestIsMaxHeapFalse(t *testing.T) {
	var maxHeaps = [][]int{
		[]int{6, 1, 7},
		[]int{7, 5, 2, 4, 3, 2, 3},
	}
	for _, heap := range maxHeaps {
		if IsMaxHeap(heap) {
			t.Error(heap, "should return false")
		}
	}
}
