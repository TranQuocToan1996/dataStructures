package heap

import "algorithms/GenericDataStructuresAlgorithmsInGo/model"

type PriorityQueue[T model.Ordered] struct {
	infoHeap Heap[T]
}

func (PriQueue *PriorityQueue[T]) Push(item T) {
	PriQueue.infoHeap.Insert(item)
}

func (PriQueue *PriorityQueue[T]) Pop() T {
	largest := PriQueue.infoHeap.Largest()
	PriQueue.infoHeap.Remove()
	return largest
}
