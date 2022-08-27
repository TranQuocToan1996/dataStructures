package c6queue

type Node[T any] struct {
	item T
	next *Node[T]
}

type NodeQueue[T any] struct {
	first, last *Node[T]
	length      int
}

type NodeIterator[T any] struct {
	next *Node[T]
}

func (q *NodeQueue[T]) Insert(item T) {
	newNode := &Node[T]{item, nil}
	if q.first == nil {
		q.first = newNode
		q.last = q.first
	} else {
		q.last.next = newNode
		q.last = newNode
	}
}

func (queue *NodeQueue[T]) Remove() T {
	returnValue := queue.first.item
	queue.first = queue.first.next
	if queue.first == nil {
		queue.last = nil
	}
	return returnValue
}

func (queue NodeQueue[T]) First() T {
	return queue.first.item
}

func (queue NodeQueue[T]) Size() int {
	return queue.length
}

func (queue *NodeQueue[T]) Range() NodeIterator[T] {
	return NodeIterator[T]{queue.first}
}
func (iterator *NodeIterator[T]) Empty() bool {
	return iterator.next == nil
}
func (iterator *NodeIterator[T]) Next() T {
	returnValue := iterator.next.item
	if iterator.next != nil {
		iterator.next = iterator.next.next
	}
	return returnValue
}
