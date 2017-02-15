package utils

import "math/rand"

// GetIntArray returns a array of pseudo-random non-negative integers of given size
func GetIntArray(l int) []int {
	if l <= 0 {
		return make([]int, 0)
	}
	seed := int64(666)
	r := rand.New(rand.NewSource(seed))
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = r.Intn(10000)
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
