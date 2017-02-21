package sorting

// QuickSort Best: O(nlgn), Average: O(nlgn), Worst: O(n^2), Space: O(n)
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, lo, hi int) {
	if lo < hi {
		pivot := partition(arr, lo, hi)
		quickSort(arr, lo, pivot-1)
		quickSort(arr, pivot+1, hi)
	}
}

// Lomuto partition scheme
func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[hi] = arr[hi], arr[i]
	return i
}
