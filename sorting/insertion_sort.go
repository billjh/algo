package sorting

// InsertionSort Best: O(n), Average: O(n^2), Worst: O(n^2), Space: O(1)
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}
