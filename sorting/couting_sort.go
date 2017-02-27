package sorting

// CountingSort Time: O(n+k), Space: O(n+k)
func CountingSort(arr []int, k int) {
	c := make([]int, k)
	for i := 0; i < len(arr); i++ {
		c[arr[i]]++
	}
	// pre-fix sum
	for i, sum := 0, 0; i < k; i++ {
		// tmp := c[i]
		// c[i] = sum
		// sum += tmp
		sum, c[i] = sum+c[i], sum
	}
	sorted := make([]int, len(arr))
	for _, n := range arr {
		sorted[c[n]] = n
		c[n]++
	}
	copy(arr, sorted)
}
