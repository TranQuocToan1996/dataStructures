package c6queue

type Passenger struct {
	name     string
	priority int
}

type PriorityQueue[T any] struct {
	q    []NodeQueue[T]
	size int
}

func NewPriorityQueue[T any](numberPriorities int) (pq PriorityQueue[T]) {
	pq.q = make([]NodeQueue[T], numberPriorities)
	return pq
}

func (pq *PriorityQueue[T]) Insert(item T, priority int) {
	pq.q[priority-1].Insert(item)
	pq.size++
}

func (pq *PriorityQueue[T]) Remove() T {
	pq.size--
	for i := 0; i < len(pq.q); i++ {
		if pq.q[i].Size() > 0 {
			return pq.q[i].Remove()
		}
	}
	var zero T
	return zero
}

func (pq *PriorityQueue[T]) First() T {
	for _, queue := range pq.q {
		if queue.Size() > 0 {
			return queue.First()
		}
	}
	var zero T
	return zero
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	result := true
	for _, queue := range pq.q {
		if queue.Size() > 0 {
			result = false
			break
		}
	}
	return result
}
