package sorting

// HeapSort Best: O(nlgn), Average: O(nlgn), Worst: O(nlgn), Space: O(1)
func HeapSort(arr []int) {
	// heapify the array into a max-heap
	heapify(arr)
	for i := len(arr) - 1; i > 0; i-- {
		// swap the root of the max-heap with the last item
		arr[0], arr[i] = arr[i], arr[0]
		// fix heap
		siftDown(arr, 0, i)
	}
}

func heapify(arr []int) {
	for i := (len(arr) - 1) / 2; i >= 0; i-- {
		siftDown(arr, i, len(arr))
	}
}

func siftDown(heap []int, lo, hi int) {
	root := lo
	for {
		child := root*2 + 1
		if child >= hi {
			break
		}
		if child+1 < hi && heap[child] < heap[child+1] {
			child++
		}
		if heap[root] < heap[child] {
			heap[root], heap[child] = heap[child], heap[root]
			root = child
		} else {
			break
		}

	}
}
