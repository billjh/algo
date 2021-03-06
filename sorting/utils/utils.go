package utils

import "math/rand"

// GetIntArray returns a array of pseudo-random non-negative integers of given size l
func GetIntArray(l int) []int {
	if l <= 0 {
		return make([]int, 0)
	}
	seed := int64(666)
	r := rand.New(rand.NewSource(seed))
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = r.Int()
	}
	return arr
}

// GetIntNArray returns a array of pseudo-random 0~n integers of given size l
func GetIntNArray(l, n int) []int {
	if l <= 0 {
		return make([]int, 0)
	}
	seed := int64(666)
	r := rand.New(rand.NewSource(seed))
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = r.Intn(n)
	}
	return arr
}

// IsSorted checks if the given integer array is sorted
func IsSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// IsMaxHeap checks if the given array is a max-heap (binary heap)
func IsMaxHeap(heap []int) bool {
	l := len(heap)
	for i := 0; i < l; i++ {
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		if leftChild < l && heap[i] < heap[leftChild] {
			return false
		}
		if rightChild < l && heap[i] < heap[rightChild] {
			return false
		}
	}
	return true
}
