package sorting

func MergeSort(arr []int) {
	copy(arr, mergeSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := mergeSort(arr[:middle])
	right := mergeSort(arr[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	r := make([]int, len(left)+len(right))
	for i := 0; ; i++ {
		if len(left) > 0 && len(right) > 0 {
			// pick one smaller item when two arrays are both non-empty
			if left[0] > right[0] {
				r[i] = right[0]
				right = right[1:]
			} else {
				r[i] = left[0]
				left = left[1:]
			}
		} else if len(left) > 0 {
			copy(r[i:], left)
			break
		} else if len(right) > 0 {
			copy(r[i:], right)
			break
		}
	}
	return r
}
