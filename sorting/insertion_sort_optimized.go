package sorting

func OptimizedInsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		unsorted := arr[i]
		j := i - 1
		for ; j >= 0 && unsorted < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = unsorted
	}
}
