package sorting

// SelectionSort Time: Θ(n^2), Space: O(1)
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		iMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[iMin] {
				iMin = j
			}
		}
		if iMin != i {
			arr[i], arr[iMin] = arr[iMin], arr[i]
		}
	}
}
