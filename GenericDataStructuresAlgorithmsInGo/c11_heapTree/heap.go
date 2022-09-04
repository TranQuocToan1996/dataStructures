package heap

import "algorithms/GenericDataStructuresAlgorithmsInGo/model"

// go lib: https://pkg.go.dev/container/heap

// A heap tree is
// another balanced tree type with the largest item in the tree always in the root of the tree.
// We use a heap tree to implement an efficient sorting algorithm

// A heap is a complete binary tree such that each node has a value greater than its two children

// The largest value in a heap tree will always be in the root node. A complete tree
// has leaf nodes filled from left to right, all at the deepest level in the tree

// We can only delete the value in the root node of a heap tree

// we replace the value in the root node with the value in the rightmost node on the lowest level of the tree

type Heap[T model.Ordered] struct {
	Items []T
}

func (h *Heap[T]) Swap(index1, index2 int) {

	if len(h.Items) < 2 {
		return
	}

	h.Items[index1], h.Items[index2] =
		h.Items[index2], h.Items[index1]
}

func NewHeap[T model.Ordered](input []T) *Heap[T] {
	heap := &Heap[T]{}
	for i := 0; i < len(input); i++ {
		heap.Insert(input[i])
	}
	return heap
}

func (heap *Heap[T]) Insert(value T) {
	heap.Items = append(heap.Items, value)
	heap.buildHeap(len(heap.Items) - 1)
}

func (heap *Heap[T]) Remove() {
	if len(heap.Items) == 0 {
		return
	}
	// Can only remove Items[0], the largest value
	heap.Items[0] = heap.Items[len(heap.Items)-1]
	heap.Items = heap.Items[:(len(heap.Items) - 1)]
	heap.rebuildHeap(0)
}

func (heap *Heap[T]) Largest() T {
	if len(heap.Items) == 0 {
		var res T
		return res
	}
	return heap.Items[0]
}

// works upward from the bottom of the tree doing swaps when necessary to produce a heap
func (heap *Heap[T]) buildHeap(index int) {
	var parent int
	if index > 0 {
		parent = (index - 1) / 2
		if heap.Items[index] > heap.Items[parent] {
			heap.Swap(index, parent)
		}
		heap.buildHeap(parent)
	}
}

// rebuildHeap compares the new root value with the
// values of its two children, swapping with the larger of the children
// continue this recursive process until there are no further nodes to swap
func (heap *Heap[T]) rebuildHeap(index int) {
	length := len(heap.Items)
	if (2*index + 1) < length {
		left := 2*index + 1
		right := 2*index + 2
		largest := index
		if left < length && right < length &&
			heap.Items[left] >= heap.Items[right] &&
			heap.Items[index] < heap.Items[left] {
			largest = left
		} else if right < length &&
			heap.Items[right] >= heap.Items[left] &&
			heap.Items[index] < heap.Items[right] {
			largest = right
		} else if left < length && right >= length &&
			heap.Items[index] < heap.Items[left] {
			largest = left
		}
		if index != largest {
			heap.Swap(index, largest)
			heap.rebuildHeap(largest)
		}
	}
}
