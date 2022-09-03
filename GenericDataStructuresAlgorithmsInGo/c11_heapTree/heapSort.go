package heap

import "algorithms/GenericDataStructuresAlgorithmsInGo/model"

// Build a heap from the initial list to be sorted. Extract the largest from the root and
// append it to the result list (initialized to empty). Apply the Remove method to the heap.
// This process will produce a slice sorted from largest to smallest (descending)

func HeapSort[T model.Ordered](data []T, isAscending bool) []T {
	heap := NewHeap[T](data)
	descending := []T{}
	for len(heap.Items) > 0 {
		descending = append(descending, heap.Largest())
		heap.Remove()
	}

	if isAscending {
		endIndex := len(descending) - 1
		ascending := make([]T, len(data))
		for i := endIndex; i >= 0; i-- {
			ascending[endIndex-i] = descending[i]
		}
		return ascending
	}

	return descending
}

func IsSorted[T model.Ordered](data []T) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}
