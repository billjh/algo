package sorting

// BubbleSort Best: O(n), Average: O(n^2), Worst: O(n^2), Space: O(1)
func BubbleSort(arr []int) {
	n := len(arr)
	for n > 0 {
		newn := 0
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				newn = i
			}
		}
		n = newn
	}
}
