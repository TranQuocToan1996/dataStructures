package c6queue

type Queue[T any] struct {
	items []T
}

type Iterator[T any] struct {
	next  int // index in items
	items []T
}

func (q *Queue[T]) Insert(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) First() T {
	var resp T
	if len(q.items) > 0 {
		resp = q.items[0]
	}
	return resp
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Range() *Iterator[T] {
	return &Iterator[T]{0, q.items}
}

func (queue *Queue[T]) Remove() T {
	returnValue := queue.items[0]
	queue.items = queue.items[1:]
	return returnValue
}

func (i *Iterator[T]) Empty() bool {
	return i.next == len(i.items)
}

func (i *Iterator[T]) Next() T {
	next := i.items[i.next]
	i.next++
	return next
}
